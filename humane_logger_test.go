package sklog

import (
	"bytes"
	"testing"
	"time"

	"net/http"

	"github.com/stretchr/testify/assert"
)

func TestHumaneLogger_Log(t *testing.T) {
	b := bytes.NewBuffer(nil)
	e := bytes.NewBuffer(nil)
	l := NewHumaneLogger(b, DefaultHTTPFormatter)
	n := time.Now()

	err := l.Log(
		KeyMessage, "log message",
		KeyTimestamp, n.Format(time.RFC3339),
		KeyHTTPMethod, "GET",
		KeyLevel, LevelDebug,
		KeySubsystem, "creative-service",
		KeyHTTPStatus, http.StatusInternalServerError,
		"field1", "value1",
	)

	NewKeyFormatter(formatBraces, KeyTimestamp).Format(e, n.Format(time.RFC3339))
	NewKeyFormatter(formatBracesLevel, KeyLevel).Format(e, LevelDebug)
	NewKeyFormatter(formatBraces, KeySubsystem).Format(e, "creative-service")
	NewKeyFormatter(formatBraces, KeyHTTPMethod).Format(e, "GET")
	NewKeyFormatter(formatBraces, KeyHTTPStatus).Format(e, http.StatusInternalServerError)
	NewKeyFormatter(formatMessage, KeyMessage).Format(e, "log message")
	e.WriteString("field1=value1  \n")

	if assert.NoError(t, err) {
		assert.Equal(t, e.String(), b.String())
	}
}
