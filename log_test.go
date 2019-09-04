package test

import (
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	testLog := Log{}
	l := log.New(&testLog, "test ", 0)
	l.Println("test line")
	testLog.Fixture(t)
	testLog.Contains(t, "line")
	testLog.DoesntContain(t, "Line")
}
