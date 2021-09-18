package test

import (
	"bytes"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
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

func (l *Log) Get() []byte {
	l.Lock()
	defer l.Unlock()

	return l.buf.Bytes()
}

func (l *Log) Fixture(t *testing.T) {
	t.Helper()

	l.Lock()
	defer l.Unlock()

	Fixture(t, l.buf.String())
}

func (l *Log) Contains(t *testing.T, substr string) {
	t.Helper()

	l.Lock()
	defer l.Unlock()

	require.True(t, strings.Contains(l.buf.String(), substr))
}

func (l *Log) DoesntContain(t *testing.T, substr string) {
	t.Helper()

	l.Lock()
	defer l.Unlock()

	require.False(t, strings.Contains(l.buf.String(), substr))
}
