package main

import (
	"cmp"
	"html/template"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

func eachRom(root string, cb func(romFile string)) error {
	dir, err := os.ReadDir(root)

	if err != nil {
		return err
	}

	for _, entry := range dir {
		name := entry.Name()
		fullPath := filepath.Join(root, name)

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
			{{- if .SerialCallbackFuncName }}
			{{ .SerialCallbackFuncName }}(),
			{{ end -}}
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
	"log/slog"
	"testing"
)

func Test{{ .Name }}Roms(t *testing.T) {
	cleanupOutputs(t)
	t.Parallel()

	slog.SetLogLoggerLevel(slog.LevelDebug)
	{{ template "subtest" .Roms }}
}
`))

var benchTemplate = template.Must(template.New("benchRoms").
	Parse(`
{{- define "child_tests" }}
{{- range .Children -}}
	{{ template "subtest" . }}
{{- end -}}
{{ end -}}
{{- define "subtest" }}
	{{ if .Name }}
	b.Run("{{ .Name }}", func(b *testing.B) {
		{{- if .FullPath -}}
		runRomBenchmark(b, []serialOutCallbackCreator{
			{{- if .SerialCallbackFuncName }}
			{{ .SerialCallbackFuncName }},
			{{ end -}}
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

func Benchmark{{ .Name }}Roms(b *testing.B) {
	{{ template "subtest" .Roms }}
}
`))

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

var romTestCollections = []struct {
	Name                   string
	RomsRoot               string
	SerialCallbackFuncName string
	OutFile                string
}{
	{"Blargg", "../../testdata/roms/blargg/", "blarggSerialCallback", "gb_roms_blargg_test.go"},
	{"Mooneye", "../../testdata/roms/mooneye/", "mooneyeSerialCallback", "gb_roms_mooneye_test.go"},
	{"Misc", "../../testdata/roms/misc/", "", "gb_roms_misc_test.go"},
}

var romBenchCollections = []struct {
	Name                   string
	RomsRoot               string
	SerialCallbackFuncName string
	OutFile                string
}{
	{"BlarggCpuInstrs", "../../testdata/roms/blargg/cpu_instrs", "blarggSerialCallback", "gb_roms_blargg_bench_test.go"},
	{"MooneyeMBC1", "../../testdata/roms/mooneye/emulator-only/mbc1", "mooneyeSerialCallback", "gb_roms_mooneye_bench_test.go"},
}

func main() {
	for _, rc := range romTestCollections {
		f, err := os.Create(rc.OutFile)

		if err != nil {
			slog.Error(err.Error())
			continue
		}

		defer func(f *os.File) {
			err := f.Close()

			if err != nil {
				slog.Error(err.Error())
			}
		}(f)

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

	for _, rc := range romBenchCollections {
		f, err := os.Create(rc.OutFile)

		if err != nil {
			slog.Error(err.Error())
			continue
		}

		defer func(f *os.File) {
			err := f.Close()

			if err != nil {
				slog.Error(err.Error())
			}
		}(f)

		tree, err := getTestCaseTree(rc.RomsRoot, rc.SerialCallbackFuncName)

		if err != nil {
			slog.Error(err.Error())
			continue
		}

		err = benchTemplate.Execute(f, struct {
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

		slog.Info("generated benchmarks for roms", "collectionName", rc.Name)
	}
}
