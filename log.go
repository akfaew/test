package test

import (
	"bytes"
	"strings"
	"sync"
	"testing"
)

type Log struct {
	sync.Mutex
	buf bytes.Buffer
}

func (l *Log) Write(data []byte) (n int, err error) {
	l.Lock()
	defer l.Unlock()

	return l.buf.Write(data)
}

func (l *Log) Reset() {
	l.Lock()
	defer l.Unlock()

	l.buf.Reset()
}

func (l *Log) Bytes() []byte {
	l.Lock()
	defer l.Unlock()

	return l.buf.Bytes()
}

func (l *Log) String() string {
	l.Lock()
	defer l.Unlock()

	return l.buf.String()
}

func (l *Log) Fixture(t *testing.T) {
	t.Helper()

	l.Lock()
	defer l.Unlock()

	FixtureExtra(t, "log", l.buf.String())
}

func (l *Log) Contains(t *testing.T, substr string) {
	t.Helper()

	l.Lock()
	defer l.Unlock()

	True(t, strings.Contains(l.buf.String(), substr))
}

func (l *Log) DoesntContain(t *testing.T, substr string) {
	t.Helper()

	l.Lock()
	defer l.Unlock()

	False(t, strings.Contains(l.buf.String(), substr))
}
