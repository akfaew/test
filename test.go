// Package test provides convenience functions for testing
//
// Fixtures are written to "testdata/output/" and read from "testdata/input/", configurable with
// OutputPath and InputPath respectively.
package test

import (
	"reflect"
	"testing"
)

// True ensures that val is true
func True(t *testing.T, val bool) {
	t.Helper()

	if val != true {
		t.Fatalf("Condition is %v, but it should be true", val)
	}
}

// False ensures that val is false
func False(t *testing.T, val bool) {
	t.Helper()

	if val != false {
		t.Fatalf("Condition is %v, but it should be false", val)
	}
}

// DeepEqual ensures that got DeepEquals expected
func DeepEqual(t *testing.T, got, expected interface{}) {
	t.Helper()

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("\n"+
			"got:      [%+v]\n"+
			"expected: [%+v]", got, expected)
	}
}

// NoError ensures that err is nil
func NoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("err=%+v", err)
	}
}

// Len ensures that len(arr) equals l
func Len(t *testing.T, arr interface{}, l int) {
	t.Helper()

	arrV := reflect.ValueOf(arr)
	if arrV.Kind() != reflect.Slice {
		t.Fatalf("Second argument must be a slice, and not %v", arrV.Kind())
	}

	if arrV.Len() != l {
		t.Fatalf("Length of array is %v, but it should be %v", arrV.Len(), l)
	}
}
