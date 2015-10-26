package sklog_test

import (
	"bytes"
	"testing"
	"time"

	"net/http"

	"github.com/piotrkowalczuk/sklog"
	"github.com/stretchr/testify/assert"
)

func TestHumaneLogger_Log(t *testing.T) {
	b := bytes.NewBuffer(nil)
	e := bytes.NewBuffer(nil)
	l := sklog.NewHumaneLogger(b)
	n := time.Now()

	err := l.Log(
		sklog.KeyMessage, "log message",
		sklog.KeyTimestamp, n.Format(time.RFC3339),
		sklog.KeyHTTPMethod, "GET",
		sklog.KeyLevel, sklog.LevelDebug,
		sklog.KeySubsystem, "creative-service",
		sklog.KeyHTTPStatus, http.StatusInternalServerError,
		"field1", "value1",
	)

	sklog.NewBracesFormatter(sklog.KeyTimestamp, 0).Format(e, n.Format(time.RFC3339))
	sklog.NewBracesFormatter(sklog.KeyLevel, 5).Format(e, sklog.LevelDebug)
	sklog.NewBracesFormatter(sklog.KeySubsystem, 0).Format(e, "creative-service")
	sklog.NewBracesFormatter(sklog.KeyHTTPMethod, 0).Format(e, "GET")
	sklog.NewBracesFormatter(sklog.KeyHTTPStatus, 0).Format(e, http.StatusInternalServerError)
	sklog.NewMessageFormatter(sklog.KeyMessage).Format(e, "log message")
	e.WriteString("field1=value1  \n")

	if assert.NoError(t, err) {
		assert.Equal(t, e.String(), b.String())
	}
}
