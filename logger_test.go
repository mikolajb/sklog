package sklog_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/piotrkowalczuk/sklog"
	"github.com/stretchr/testify/assert"
)

func testLevel(t *testing.T, scenarios ...func(*testing.T, *bytes.Buffer, log.Logger)) {
	b := bytes.NewBuffer(nil)
	l := log.NewJSONLogger(b)

	for _, scenario := range scenarios {
		scenario(t, b, l)
		b.Reset()
	}
}
func TestSetTimestampFunc(t *testing.T) {
	testLevel(t, testSetTimestampFunc)
}

func testSetTimestampFunc(t *testing.T, b *bytes.Buffer, l log.Logger) {
	fn := func() string { return "fake-timestamp" }
	sklog.SetTimestampFunc(fn)
	sklog.Info(l, "TEST")

	assert.Contains(t, b.String(), `"msg":"TEST"`)
	assert.Contains(t, b.String(), `"timestamp":"`+fn()+`"`)
}

func TestInfo(t *testing.T) {
	testLevel(t, testInfoWithOnlyMessage, testInfoWithMessageAndTag)
}

func testInfoWithOnlyMessage(t *testing.T, b *bytes.Buffer, l log.Logger) {
	sklog.Info(l, "TEST")

	assert.Contains(t, b.String(), `"level":"info"`)
	assert.Contains(t, b.String(), `"msg":"TEST"`)
	assert.Contains(t, b.String(), `"timestamp":`)
}

func testInfoWithMessageAndTag(t *testing.T, b *bytes.Buffer, l log.Logger) {
	sklog.Info(l, "TEST", "tag1", "value1")

	assert.Contains(t, b.String(), `"level":"info"`)
	assert.Contains(t, b.String(), `"msg":"TEST"`)
	assert.Contains(t, b.String(), `"timestamp":`)
	assert.Contains(t, b.String(), `"tag1":"value1"`)
}

func TestDebug(t *testing.T) {
	testLevel(t, testDebugWithOnlyMessage, testDebugWithMessageAndTag)
}

func testDebugWithOnlyMessage(t *testing.T, b *bytes.Buffer, l log.Logger) {
	sklog.Debug(l, "TEST")

	assert.Contains(t, b.String(), `"level":"debug"`)
	assert.Contains(t, b.String(), `"msg":"TEST"`)
	assert.Contains(t, b.String(), `"timestamp":`)
}

func testDebugWithMessageAndTag(t *testing.T, b *bytes.Buffer, l log.Logger) {
	sklog.Debug(l, "TEST", "tag1", "value1")

	assert.Contains(t, b.String(), `"level":"debug"`)
	assert.Contains(t, b.String(), `"msg":"TEST"`)
	assert.Contains(t, b.String(), `"timestamp":`)
	assert.Contains(t, b.String(), `"tag1":"value1"`)
}

func TestError(t *testing.T) {
	testLevel(t, testErrorWithOnlyMessage, testErrorWithMessageAndTag)
}

func testErrorWithOnlyMessage(t *testing.T, b *bytes.Buffer, l log.Logger) {
	err := errors.New("sklog_test: example error")
	sklog.Error(l, err)

	assert.Contains(t, b.String(), `"level":"error"`)
	assert.Contains(t, b.String(), `"msg":"`+err.Error()+`"`)
	assert.Contains(t, b.String(), `"timestamp":`)
}

func testErrorWithMessageAndTag(t *testing.T, b *bytes.Buffer, l log.Logger) {
	err := errors.New("sklog_test: example error")
	sklog.Error(l, err, "tag1", "value1")

	assert.Contains(t, b.String(), `"level":"error"`)
	assert.Contains(t, b.String(), `"msg":"`+err.Error()+`"`)
	assert.Contains(t, b.String(), `"timestamp":`)
	assert.Contains(t, b.String(), `"tag1":"value1"`)
}

func TestPanic(t *testing.T) {
	testLevel(t, testPanicWithOnlyMessage, testPanicWithMessageAndTag)
}

func testPanicWithOnlyMessage(t *testing.T, b *bytes.Buffer, l log.Logger) {
	err := errors.New("sklog_test: example fatal error")
	assert.Panics(t, func() {
		sklog.Panic(l, err)
	})

	assert.Contains(t, b.String(), `"level":"panic"`)
	assert.Contains(t, b.String(), `"msg":"`+err.Error()+`"`)
	assert.Contains(t, b.String(), `"timestamp":`)
}

func testPanicWithMessageAndTag(t *testing.T, b *bytes.Buffer, l log.Logger) {
	err := errors.New("sklog_test: example fatal error")
	assert.Panics(t, func() {
		sklog.Panic(l, err, "tag1", "value1")
	})

	assert.Contains(t, b.String(), `"level":"panic"`)
	assert.Contains(t, b.String(), `"msg":"`+err.Error()+`"`)
	assert.Contains(t, b.String(), `"timestamp":`)
	assert.Contains(t, b.String(), `"tag1":"value1"`)
}
