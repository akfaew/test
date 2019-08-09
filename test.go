// Package test provides convenience functions for testing
//
// Fixtures are written to "testdata/output/" and read from "testdata/input/", configurable with
// FixtureOutputPath and FixtureInputPath respectively.
package test

import (
	"reflect"
	"testing"
)

// True ensures that val is true
func True(t *testing.T, val bool) {
	t.Helper()

	if !val {
		t.Fatalf("Condition is %v, but it should be true", val)
	}
}

// False ensures that val is false
func False(t *testing.T, val bool) {
	t.Helper()

	if val {
		t.Fatalf("Condition is %v, but it should be false", val)
	}
}

// EqualStr ensures that two strings are equal
func EqualStr(t *testing.T, got, expected string) {
	t.Helper()

	if got != expected {
		t.Fatalf("Got: %q, expected: %q", got, expected)
	}
}

// EqualInt ensures that two ints are equal
func EqualInt(t *testing.T, got, expected int) {
	t.Helper()

	if got != expected {
		t.Fatalf("Got: \"%d\", expected: \"%d\"", got, expected)
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

// IsNil ensures that val is nil
func IsNil(t *testing.T, val interface{}) {
	t.Helper()

	if val != nil {
		t.Fatalf("val=%+v", val)
	}
}

// IsNotNil ensures that val is not nil
func IsNotNil(t *testing.T, val interface{}) {
	t.Helper()

	if val == nil {
		t.Fatalf("val is nil")
	}
}

// NoError ensures that err is nil
func NoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("err=%+v", err)
	}
}

// Len ensures that len(obj) equals l
func Len(t *testing.T, obj interface{}, l int) {
	t.Helper()

	objV := reflect.ValueOf(obj)
	if objV.Kind() == reflect.Slice {
		if objV.Len() != l {
			t.Fatalf("Length of array is %v, but it should be %v", objV.Len(), l)
		}
	} else if objV.Kind() == reflect.Map {
		if objV.Len() != l {
			t.Fatalf("Length of map is %v, but it should be %v", objV.Len(), l)
		}
	} else {
		t.Fatalf("Second argument must be a slice or a map, and not %v", objV.Kind())
	}
}
