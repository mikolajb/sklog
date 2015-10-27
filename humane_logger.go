package sklog

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"

	"github.com/go-kit/kit/log"
)

type humaneLogger struct {
	io.Writer
	formatters []Formatter
}

// NewHumaneLogger returns a Logger that encodes keyvals to the Writer in a human friendly way.
func NewHumaneLogger(writer io.Writer) log.Logger {
	return NewHumaneLoggerWithFormatters(writer, []Formatter{
		NewBracesFormatter(KeyTimestamp, 0),
		NewBracesFormatter(KeyLevel, 5),
		NewBracesFormatter(KeySubsystem, 0),
		NewBracesFormatter(KeyHTTPMethod, 0),
		NewBracesFormatter(KeyHTTPStatus, 0),
		NewMessageFormatter(KeyMessage),
	})
}

// NewHumaneLoggerWithFormatters like NewHumaneLogger allocates new instance,
// but allow to pass custom collection of formatters.
func NewHumaneLoggerWithFormatters(writer io.Writer, formatters []Formatter) log.Logger {
	return &humaneLogger{
		Writer:     writer,
		formatters: formatters,
	}
}

// Log implements Logger interface.
func (hl *humaneLogger) Log(keyvals ...interface{}) (err error) {
	n := (len(keyvals) + 1) / 2 // +1 to handle case when len is odd
	m := make(map[string]interface{}, n)
	b := bytes.NewBuffer(nil)

	for i := 0; i < len(keyvals); i += 2 {
		k := keyvals[i]
		var v interface{} = log.ErrMissingValue
		if i+1 < len(keyvals) {
			v = keyvals[i+1]
		}
		merge(m, k, v)
	}

	for _, formatter := range hl.formatters {
	MapLoop:
		for key, value := range m {
			if formatter.Key() == key {
				formatter.Format(b, value)
				delete(m, key)
				break MapLoop
			}
		}
	}

	for key, value := range m {
		_, err = fmt.Fprintf(b, "%s=%v  ", key, value)
		if err != nil {
			return
		}
	}

	b.WriteRune('\n')
	_, err = b.WriteTo(hl.Writer)

	return
}

// Formatter ...
type Formatter interface {
	Key() string
	Format(writer io.Writer, value interface{}) (int, error)
}

type formatter struct {
	key      string
	function func(io.Writer, interface{}) (int, error)
}

func (f *formatter) Key() string {
	return f.key
}

func (f *formatter) Format(w io.Writer, v interface{}) (int, error) {
	return f.function(w, v)
}

// NewBracesFormatter writes value for given key with surrounding braces.
func NewBracesFormatter(key string, length int) Formatter {
	format := "[%v] "

	if length > 0 {
		format = "[%-" + strconv.FormatInt(int64(length), 10) + "v] "
	}

	return &formatter{
		key: key,
		function: func(w io.Writer, value interface{}) (int, error) {
			return fmt.Fprintf(w, format, value)
		},
	}
}

// NewMessageFormatter writes value for given key prefixed with minus.
func NewMessageFormatter(key string) Formatter {
	format := "- %-60v "
	return &formatter{
		key: key,
		function: func(w io.Writer, value interface{}) (int, error) {
			return fmt.Fprintf(w, format, value)
		},
	}
}

func merge(dst map[string]interface{}, k, v interface{}) {
	var key string
	switch x := k.(type) {
	case string:
		key = x
	case fmt.Stringer:
		key = safeString(x)
	default:
		key = fmt.Sprint(x)
	}
	if x, ok := v.(error); ok {
		v = safeError(x)
	}
	dst[key] = v
}

func safeString(str fmt.Stringer) (s string) {
	defer func() {
		if panicVal := recover(); panicVal != nil {
			if v := reflect.ValueOf(str); v.Kind() == reflect.Ptr && v.IsNil() {
				s = "NULL"
			} else {
				panic(panicVal)
			}
		}
	}()
	s = str.String()
	return
}

func safeError(err error) (s interface{}) {
	defer func() {
		if panicVal := recover(); panicVal != nil {
			if v := reflect.ValueOf(err); v.Kind() == reflect.Ptr && v.IsNil() {
				s = nil
			} else {
				panic(panicVal)
			}
		}
	}()
	s = err.Error()
	return
}
