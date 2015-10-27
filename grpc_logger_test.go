package sklog_test

import (
	"bytes"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/piotrkowalczuk/sklog"
	"github.com/stretchr/testify/assert"
)

func TestGRPCLogger_Print(t *testing.T) {
	b := bytes.NewBuffer(nil)
	l := log.NewJSONLogger(b)
	g := sklog.NewGRPCLogger(l)

	success := map[string][]interface{}{
		"message": {
			"message",
		},
		"message1, message2": {
			"message1",
			"message2",
		},
		"1, [109 101 115 115 97 103 101 50]": {
			1,
			[]byte("message2"),
		},
		"true, [message1 message2]": {
			true,
			[]string{"message1", "message2"},
		},
	}

	for expected, args := range success {
		g.Print(args...)
		assert.Contains(t, b.String(), sklog.LevelDebug)
		assert.Contains(t, b.String(), sklog.KeyLevel)
		assert.Contains(t, b.String(), expected)
		b.Reset()
	}
}

func TestGRPCLogger_Printf(t *testing.T) {
	b := bytes.NewBuffer(nil)
	l := log.NewJSONLogger(b)
	g := sklog.NewGRPCLogger(l)

	success := []struct {
		expected string
		format   string
		args     []interface{}
	}{
		{
			expected: "message",
			format:   "message",
		},
		{
			expected: "message1, message2",
			format:   "%s, %s",
			args: []interface{}{
				"message1",
				"message2",
			},
		},
	}

	for _, data := range success {
		g.Printf(data.format, data.args...)
		assert.Contains(t, b.String(), sklog.LevelDebug)
		assert.Contains(t, b.String(), sklog.KeyLevel)
		assert.Contains(t, b.String(), data.expected)
		b.Reset()
	}
}
