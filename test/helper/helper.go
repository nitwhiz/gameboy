package helper

import (
	"encoding/json"
	"reflect"
	"testing"
)

func RequireJsonEqual(t *testing.T, expected string, actual []byte) {
	var e map[string]any
	var a map[string]any

	if err := json.Unmarshal([]byte(expected), &e); err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(actual, &a); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(e, a) {
		t.Errorf("json does not match.\nExpected: %+v\nActual:   %+v", e, a)
	}
}
