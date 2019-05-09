package test

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	permissions = 0644
)

var (
	regen = flag.Bool("regen", false, "Regenerate fixtures")

	FixtureInputPath  = "testdata/input/"
	FixtureOutputPath = "testdata/output/"
)

// makeFixturePath makes a path from the test name, and optionally appends "extra".
func makeFixturePath(t *testing.T, extra string) string {
	t.Helper()

	name := strings.Replace(t.Name(), "/", "-", -1)
	path := FixtureOutputPath + name
	if extra != "" {
		path += "-" + extra
	}
	path += ".fixture"

	return path
}

// Fixture ensures that 'data' is equal to what's stored on disk.
//
// If 'data' is a string it gets written verbatim, otherwise it's json-encoded.
//
// The filename of the fixture is generated from the test name. To use multiple fixtures in one test see FixtureExtra()
func Fixture(t *testing.T, data interface{}) {
	t.Helper()

	FixtureExtra(t, "", data)
}

// FixtureExtra ensures that data is equal to what's stored on disk.
//
// If 'data' is a string it gets written verbatim, otherwise it's json-encoded.
//
// The filename of the fixture is generated from the test name with 'extra' appended.
func FixtureExtra(t *testing.T, extra string, data interface{}) {
	t.Helper()

	// Write strings verbatim, otherwise json-encode.
	var got []byte
	if b, ok := data.(string); ok {
		got = []byte(b)
	} else {
		var err error
		got, err = json.Marshal(data)
		NoError(t, err)
	}

	path := makeFixturePath(t, extra)
	// If -regen then write and return
	if *regen {
		if err := ioutil.WriteFile(path, []byte(got), permissions); err != nil {
			t.Fatalf("Error writing file %q: %v", path, err)
		}
		return
	}

	want, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("Error reading file %q: %v", path, err)
	}

	if !bytes.Equal(got, want) {
		if err := ioutil.WriteFile("/tmp/got", got, permissions); err != nil {
			t.Fatalf("Error writing file /tmp/got: %v", err)
		}
		if err := ioutil.WriteFile("/tmp/want", want, permissions); err != nil {
			t.Fatalf("Error writing file /tmp/want: %v", err)
		}
		t.Fatalf("Error comparing with fixture. See: diff /tmp/got /tmp/want")
	}
}

// FixtureHTTP is like Fixture, except it reads the data from 'res' and treats it as a string type.
func FixtureHTTP(t *testing.T, res *httptest.ResponseRecorder) {
	t.Helper()

	FixtureHTTPExtra(t, "", res)
}

// FixtureHTTPExtra is like FixtureExtra, except it reads the data from 'res' and treats it as a string type.
func FixtureHTTPExtra(t *testing.T, extra string, res *httptest.ResponseRecorder) {
	t.Helper()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error reading from ResponseRecorder: %v", err)
	}

	code := fmt.Sprintf("Code: %d\n\n", res.Code)
	FixtureExtra(t, extra, code+string(body))
}

// InputFixture returns the contents of a fixture file
func InputFixture(t *testing.T, filename string) []byte {
	t.Helper()

	input, err := ioutil.ReadFile(FixtureInputPath + filename)
	if err != nil {
		t.Fatalf("Error reading fixture: %v", err)
	}

	return input
}
