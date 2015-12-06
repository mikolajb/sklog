package sklog_test

import (
	"testing"

	"github.com/piotrkowalczuk/sklog"
)

func TestTestLogger_Log(t *testing.T) {
	sklog.NewTestLogger(t).Log(sklog.KeyMessage, "I'm a test logger!")
	sklog.NewTestLogger(t).Log(sklog.KeyMessage, "I can log only messages ;(", sklog.KeyHTTPMethod, "GET")
}
