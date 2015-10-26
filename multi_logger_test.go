package sklog_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/piotrkowalczuk/sklog"
	"github.com/stretchr/testify/assert"
)

type faultyLogger struct{}

func (fl *faultyLogger) Log(keyvals ...interface{}) (err error) {
	return errors.New("sklog_test: faulty logger error")
}

func TestMultiLogger_Log(t *testing.T) {
	b1 := bytes.NewBuffer(nil)
	b2 := bytes.NewBuffer(nil)

	l1 := log.NewJSONLogger(b1)
	l2 := log.NewJSONLogger(b2)

	l := sklog.NewMultiLogger(l1, l2)

	sklog.Debug(l, "debug message")
	sklog.Info(l, "info message")
	sklog.Error(l, errors.New("sklog_test: example error message"))

	assert.Equal(t, b1.String(), b2.String())

	b1.Reset()
	b2.Reset()

	l1 = &faultyLogger{}
	l2 = log.NewJSONLogger(b2)

	l = sklog.NewMultiLogger(l1, l2)

	sklog.Debug(l, "debug message")
	sklog.Info(l, "info message")
	sklog.Error(l, errors.New("sklog_test: example error message"))

	assert.Contains(t, b2.String(), "debug message")
	assert.Contains(t, b2.String(), "msg")
	assert.Contains(t, b2.String(), "info message")
	assert.Contains(t, b2.String(), "info")
	assert.Contains(t, b2.String(), "debug")
	assert.Contains(t, b2.String(), "sklog_test: example error message")

}
