package integration

import (
	"github.com/nitwhiz/gameboy/pkg/inst"
	"reflect"
	"strings"
	"testing"
)

func TestBlarggRoms(t *testing.T) {
	t.Parallel()

	inst.InitHandlers()

	serialData := map[string][]byte{}

	testRomsRecursive(t, "../../testdata/roms/blargg/", []serialOutCallbackFunc{
		func(testName string, b byte) (bool, bool) {
			s, ok := serialData[testName]

			if !ok {
				s = []byte{}
			}

			s = append(s, b)
			serialData[testName] = s

			if strings.HasSuffix(string(s), "Passed") {
				return false, true
			}

			if strings.HasSuffix(string(s), "Failed") {
				return false, false
			}

			return true, true
		},
	})
}

func TestMooneyeRoms(t *testing.T) {
	t.Parallel()

	inst.InitHandlers()

	serialData := map[string][]byte{}

	successData := []byte{3, 5, 8, 13, 21, 34}
	failData := []byte{66, 66, 66, 66, 66, 66}

	testRomsRecursive(t, "../../testdata/roms/mooneye/", []serialOutCallbackFunc{
		func(testName string, b byte) (bool, bool) {
			s, ok := serialData[testName]

			if !ok {
				s = []byte{}
			}

			s = append(s, b)
			serialData[testName] = s

			if reflect.DeepEqual(s, successData) {
				return false, true
			}

			if reflect.DeepEqual(s, failData) {
				return false, false
			}

			return true, true
		},
	})
}
