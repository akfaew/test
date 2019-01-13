package test

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_mkpath(t *testing.T) {
	path1 := mkpath(t, "")
	if path1 != "testdata/output/Test_mkpath.fixture" {
		t.Fatalf(path1)
	}

	t.Run("Sub Test", func(t *testing.T) {
		path2 := mkpath(t, "")
		if path2 != "testdata/output/Test_mkpath-Sub_Test.fixture" {
			t.Fatalf(path2)
		}
	})

	t.Run("Sub Test With Extra", func(t *testing.T) {
		path3 := mkpath(t, "extra")
		if path3 != "testdata/output/Test_mkpath-Sub_Test_With_Extra-extra.fixture" {
			t.Fatalf(path3)
		}
	})
}

func Test_Fixture(t *testing.T) {
	t.Run("bytes", func(t *testing.T) {
		b := []byte("an array of bytes")
		Fixture(t, b)
	})

	t.Run("string", func(t *testing.T) {
		b := "a string of text"
		Fixture(t, b)
	})

	t.Run("regen", func(t *testing.T) {
		b := []byte(fmt.Sprintf("%v", time.Now()))
		*regen = true
		Fixture(t, b)
		*regen = false
		Fixture(t, b)
		os.Remove(mkpath(t, ""))
	})
}
