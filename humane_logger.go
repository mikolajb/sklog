package sklog

import (
	"bytes"
	"fmt"
	"io"
	"reflect"

	"github.com/go-kit/kit/log"
)

const (
	formatMessage     = "- %-60v "
	formatBraces      = "[%v] "
	formatBracesLevel = "[%-5v] "
)

var (
	// DefaultHTTPFormatter ...
	DefaultHTTPFormatter = NewSequentialFormatter(
		NewKeyFormatter(formatBraces, KeyTimestamp),
		NewKeyFormatter(formatBracesLevel, KeyLevel),
		NewKeyFormatter(formatBraces, KeySubsystem),
		NewKeyFormatter(formatBraces, KeyHTTPMethod),
		NewKeyFormatter(formatBraces, KeyHTTPPath),
		NewKeyFormatter(formatBraces, KeyHTTPStatus),
		NewKeyFormatter(formatMessage, KeyMessage),
	)
)

type humaneLogger struct {
	io.Writer
	formatter Formatter
}

// NewHumaneLoggerWithFormatters like NewHumaneLogger allocates new instance,
// but allow to pass custom collection of formatters.
func NewHumaneLogger(writer io.Writer, formatter Formatter) log.Logger {
	return &humaneLogger{
		Writer:    writer,
		formatter: formatter,
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

	_, err = hl.formatter.Format(b, m)
	if err != nil {
		b.Reset()
		return err
	}

	b.WriteRune('\n')
	_, err = b.WriteTo(hl.Writer)

	return
}

// Formatter ...
type Formatter interface {
	Format(io.Writer, interface{}) (int, error)
}

// KeyFormatter ...
type KeyFormatter interface {
	Formatter
	Key() string
}

type keyFormatter struct {
	key      string
	function func(io.Writer, interface{}) (int, error)
}

// NewKeyFormatter writes value for given key using given format.
func NewKeyFormatter(format, key string) KeyFormatter {
	return &keyFormatter{
		key: key,
		function: func(w io.Writer, value interface{}) (int, error) {
			return fmt.Fprintf(w, format, value)
		},
	}
}

func (kf *keyFormatter) Key() string {
	return kf.key
}

func (kf *keyFormatter) Format(w io.Writer, v interface{}) (int, error) {
	return kf.function(w, v)
}

type sequentialFormatter struct {
	formatters []KeyFormatter
}

// NewSequentialFormatter ...
func NewSequentialFormatter(f ...KeyFormatter) Formatter {
	return &sequentialFormatter{formatters: f}
}

// Format implements Formatter interface.
func (sf *sequentialFormatter) Format(w io.Writer, v interface{}) (n int, err error) {
	var n1, n2 int

	m, ok := v.(map[string]interface{})
	if !ok {
		err = fmt.Errorf("sklog: sequential formatter epxects map[string]interface{} got %T", v)
		return
	}

	for _, formatter := range sf.formatters {
	MapLoop:
		for key, value := range m {
			if formatter.Key() == key {
				n1, err = formatter.Format(w, value)
				n += n1
				if err != nil {
					return
				}
				delete(m, key)
				break MapLoop
			}
		}
	}

	n2, err = writeKV(w, m)
	n += n2

	return
}

func writeKV(w io.Writer, m map[string]interface{}) (n int, err error) {
	var n1 int
	for key, value := range m {
		n1, err = fmt.Fprintf(w, "%s=%v  ", key, value)
		n += n1
		if err != nil {
			return
		}
	}

	return
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
