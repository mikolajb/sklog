package sklog

import (
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

	tl.t.Log(m[KeyMessage])

	return nil
}
