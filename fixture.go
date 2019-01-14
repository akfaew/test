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

var regen = flag.Bool("regen", false, "Regenerate fixtures")

func mkpath(t *testing.T, extra string) string {
	t.Helper()

	name := strings.Replace(t.Name(), "/", "-", -1)
	path := "testdata/output/" + name
	if extra != "" {
		path += "-" + extra
	}
	path += ".fixture"

	return path
}

func FixtureHTTP(t *testing.T, res *httptest.ResponseRecorder) {
	t.Helper()

	FixtureHTTPExtra(t, "", res)
}

func FixtureHTTPExtra(t *testing.T, extra string, res *httptest.ResponseRecorder) {
	t.Helper()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("err=%+v", err)
	}

	code := fmt.Sprintf("Code: %d\n", res.Code)
	FixtureExtra(t, extra, code+string(body))
}

// Fixture ensures that data is equal to what's stored on disk.
// For testing multiple fixtures in one test use FixtureExtra()
func Fixture(t *testing.T, data interface{}) {
	t.Helper()

	FixtureExtra(t, "", data)
}

// FixtureExtra ensures that data is equal to what's stored on disk.
// The value in extra gets appended to the filename used.
func FixtureExtra(t *testing.T, extra string, data interface{}) {
	t.Helper()

	var dataProcessed []byte
	if b, ok := data.(string); ok {
		dataProcessed = []byte(b)
	} else {
		var err error
		dataProcessed, err = json.Marshal(data)
		NoError(t, err)
	}

	path := mkpath(t, extra)
	if *regen {
		if err := ioutil.WriteFile(path, []byte(dataProcessed), 0644); err != nil {
			t.Fatalf("ioutil.WriteFile(%s): err=%+v", path, err)
		}
		return
	}

	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("ioutil.ReadFile(%s): err=%+v", path, err)
	}

	if !bytes.Equal(dataProcessed, fileContent) {
		tmp := "/tmp/" + strings.Replace(path, "/", "-", -1)
		if err := ioutil.WriteFile(tmp, dataProcessed, 0644); err != nil {
			t.Fatalf("err=%+v", err)
		}
		t.Fatalf("Error comparing with fixture.\nFixture: <%s>\nWant:    <%s>\ndiff %s %s",
			string(fileContent), dataProcessed, path, tmp)
	}
}

// InputFixture returns the contents of a fixture file
func InputFixture(t *testing.T, name string) []byte {
	t.Helper()

	input, err := ioutil.ReadFile("testdata/input/" + name)
	if err != nil {
		t.Fatalf("err=%+v", err)
	}

	return input
}
