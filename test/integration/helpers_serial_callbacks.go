package integration

import (
	"reflect"
	"strings"
)

func blarggSerialCallback() serialOutCallbackFunc {
	var serialData []byte

	return func(b byte) (bool, bool) {
		serialData = append(serialData, b)

		if strings.HasSuffix(string(serialData), "Passed") {
			return false, true
		}

		if strings.HasSuffix(string(serialData), "Failed") {
			return false, false
		}

		return true, true
	}
}

func mooneyeSerialCallback() serialOutCallbackFunc {
	var serialData []byte

	successData := []byte{3, 5, 8, 13, 21, 34}
	failData := []byte{66, 66, 66, 66, 66, 66}

	return func(b byte) (bool, bool) {
		serialData = append(serialData, b)

		if reflect.DeepEqual(serialData, successData) {
			return false, true
		}

		if reflect.DeepEqual(serialData, failData) {
			return false, false
		}

		return true, true
	}
}
