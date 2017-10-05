package sklog

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/go-kit/kit/log"
)

type testLogger struct {
	t *testing.T
}

// NewMultiLogger returns a Logger that wraps testing object.
func NewTestLogger(t *testing.T) log.Logger {
	return &testLogger{
		t: t,
	}
}

// Log implements Logger interface.
func (tl *testLogger) Log(keyvals ...interface{}) error {
	tl.t.Helper()
	n := (len(keyvals) + 1) / 2 // +1 to handle case when len is odd
	m := make(map[string]interface{}, n)

	for i := 0; i < len(keyvals); i += 2 {
		k := keyvals[i]
		var v interface{} = log.ErrMissingValue
		if i+1 < len(keyvals) {
			v = keyvals[i+1]
		}
		merge(m, k, v)
	}

	buf := bytes.NewBuffer(nil)
	if msg, ok := m[KeyMessage]; ok {
		fmt.Fprintf(buf, "%-60s", msg)
		delete(m, KeyMessage)

	}
	for k, v := range m {
		fmt.Fprintf(buf, "%s=%v  ", k, v)
	}
	tl.t.Log(buf.String())

	return nil
}
