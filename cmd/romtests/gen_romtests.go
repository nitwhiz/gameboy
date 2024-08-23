package main

import (
	"cmp"
	"html/template"
	"log/slog"
	"os"
	"path"
	"slices"
	"strings"
	"time"
)

type testCase struct {
	FullPath               string
	Name                   string
	SerialCallbackFuncName string
	Children               []*testCase
}

var testTemplate = template.Must(template.New("testRoms").
	Parse(`
{{- define "child_tests" }}
{{- range .Children -}}
	{{ template "subtest" . }}
{{- end -}}
{{ end -}}
{{- define "subtest" }}
	{{ if .Name }}
	t.Run("{{ .Name }}", func(t *testing.T) {
		t.Parallel()
		{{ if .FullPath }}
		runRomTest(t, []serialOutCallbackFunc{
			{{ .SerialCallbackFuncName }}(),
		}, "{{ .FullPath }}", context.Background())
		{{ end -}}

		{{- template "child_tests" . -}}
	})
	{{ else }}
	{{ template "child_tests" . }}
	{{ end }}
{{ end -}}

// Do not edit. This is auto-generated.
// Timestamp: {{ .Timestamp }}

package integration

import (
	"context"
	"testing"
)

func Test{{ .Name }}Roms(t *testing.T) {
	t.Parallel()
	{{ template "subtest" .Roms }}
}
`))

func eachRom(root string, cb func(romFile string)) error {
	dir, err := os.ReadDir(root)

	if err != nil {
		return err
	}

	for _, entry := range dir {
		name := entry.Name()
		fullPath := path.Join(root, name)

		if !entry.IsDir() {
			if strings.HasSuffix(name, ".gb") {
				cb(fullPath)
			}
		} else {
			if err := eachRom(fullPath, cb); err != nil {
				return err
			}
		}
	}

	return nil
}

func getTestCaseTree(root string, serialCallbackFuncName string) (*testCase, error) {
	var romPaths []string

	t := &testCase{
		FullPath: "",
		Name:     "",
		Children: []*testCase{},
	}

	err := eachRom(root, func(romFile string) {
		romPaths = append(romPaths, romFile)
	})

	if err != nil {
		return nil, err
	}

	slices.SortStableFunc(romPaths, func(a, b string) int {
		return cmp.Compare(a, b)
	})

	for _, romPath := range romPaths {
		pathSegments := strings.Split(romPath, "/")[len(strings.Split(root, "/"))-1:]
		lastSegmentIndex := len(pathSegments) - 1

		insertAt := t

		for sIdx, segment := range pathSegments {
			if sIdx == lastSegmentIndex {
				insertAt.Children = append(insertAt.Children, &testCase{
					FullPath:               romPath,
					Name:                   strings.TrimPrefix(strings.TrimSuffix(segment, ".gb"), root),
					SerialCallbackFuncName: serialCallbackFuncName,
				})

				break
			}

			found := false

			for _, i := range insertAt.Children {
				if i.Name == segment {
					insertAt = i
					found = true

					break
				}
			}

			if !found {
				newInsertAt := &testCase{
					Name:     segment,
					Children: []*testCase{},
				}

				insertAt.Children = append(insertAt.Children, newInsertAt)
				insertAt = newInsertAt
			}
		}
	}

	return t, nil
}

var romCollections = []struct {
	Name                   string
	RomsRoot               string
	SerialCallbackFuncName string
	OutFile                string
}{
	{"Blargg", "../../testdata/roms/blargg/", "blarggSerialCallback", "gb_roms_blargg_test.go"},
	{"Mooneye", "../../testdata/roms/mooneye/", "mooneyeSerialCallback", "gb_roms_mooneye_test.go"},
}

func main() {
	for _, rc := range romCollections {
		f, err := os.Create(rc.OutFile)

		if err != nil {
			slog.Error(err.Error())
			continue
		}

		defer f.Close()

		tree, err := getTestCaseTree(rc.RomsRoot, rc.SerialCallbackFuncName)

		if err != nil {
			slog.Error(err.Error())
			continue
		}

		err = testTemplate.Execute(f, struct {
			Timestamp string
			Roms      *testCase
			Name      string
		}{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Roms:      tree,
			Name:      rc.Name,
		})

		if err != nil {
			slog.Error(err.Error())
			continue
		}

		slog.Info("generated tests for roms", "collectionName", rc.Name)
	}
}
