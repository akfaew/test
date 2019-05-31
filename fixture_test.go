package test

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_makeFixturePath(t *testing.T) {
	EqualStr(t, makeFixturePath(t, ""), "testdata/output/Test_makeFixturePath.fixture")

	t.Run("Sub Test", func(t *testing.T) {
		EqualStr(t, makeFixturePath(t, ""), "testdata/output/Test_makeFixturePath-Sub_Test.fixture")
	})

	t.Run("Sub Test With Extra", func(t *testing.T) {
		EqualStr(t, makeFixturePath(t, "extra"), "testdata/output/Test_makeFixturePath-Sub_Test_With_Extra-extra.fixture")
	})
}

func Test_Fixture(t *testing.T) {
	t.Run("bytes", func(t *testing.T) {
		Fixture(t, []byte("an array of bytes"))
	})

	t.Run("string", func(t *testing.T) {
		Fixture(t, "a string of text")
	})

	t.Run("regen", func(t *testing.T) {
		b := []byte(fmt.Sprintf("%v", time.Now()))
		*regen = true
		Fixture(t, b)
		*regen = false
		Fixture(t, b)
		os.Remove(makeFixturePath(t, ""))
	})
}

func Test_InputFixture(t *testing.T) {
	input := InputFixture(t, "input.fixture")
	EqualStr(t, string(input), "foo")
}

func Test_InputFixtureJson(t *testing.T) {
	a := struct {
		A string
		B string
		C int
	}{
		"aaa", "bbb", 123,
	}

	b := struct {
		A string
		B string
		C int
	}{}
	InputFixtureJson(t, "struct.json", &b)
	DeepEqual(t, a, b)
}
