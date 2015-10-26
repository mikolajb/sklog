package sklog

import (
	"fmt"
	"os"
	"time"

	"github.com/go-kit/kit/log"
)

const (
	// LevelDebug ...
	LevelDebug = "debug"
	// LevelError ...
	LevelError = "error"
	// LevelFatal ...
	LevelFatal = "fatal"
	// LevelInfo ...
	LevelInfo = "info"
	// LevelPanic ...
	LevelPanic = "panic"
	// LevelWarning ...
	LevelWarning = "warn"

	// KeySubsystem ...
	KeySubsystem = "subsystem"
	// KeyHTTPStatus ...
	KeyHTTPStatus = "http_status"
	// KeyHTTPMethod ...
	KeyHTTPMethod = "http_method"
	// KeyTimestamp ...
	KeyTimestamp = "timestamp"
	// KeyLevel ...
	KeyLevel = "level"
	// KeyMessage ...
	KeyMessage = "msg"
)

var (
	timestampFunc = now
)

func now() string {
	return time.Now().Format(time.RFC3339)
}

// SetTimestampFunc sets function that is used to populate timestamp field.
func SetTimestampFunc(fn func() string) {
	timestampFunc = fn
}

// Debug log message and given context with level debug.
func Debug(logger log.Logger, msg string, keyval ...interface{}) {
	logger.Log(append(keyval, KeyLevel, LevelDebug, KeyMessage, msg, KeyTimestamp, timestampFunc())...)
}

// Info log message and given context with level info.
func Info(logger log.Logger, msg string, keyval ...interface{}) {
	logger.Log(append(keyval, KeyLevel, LevelInfo, KeyMessage, msg, KeyTimestamp, timestampFunc())...)
}

// Warning log message using given logger.
func Warning(logger log.Logger, msg string, keyval ...interface{}) {
	logger.Log(append(keyval, KeyLevel, LevelWarning, KeyMessage, msg, KeyTimestamp, timestampFunc())...)
}

// Error log error using given logger.
func Error(logger log.Logger, err error, keyval ...interface{}) {
	contextErrorFunc(logger, err).Log(append(keyval, KeyLevel, LevelError, KeyMessage, err.Error(), KeyTimestamp, timestampFunc())...)
}

// Fatal log error using given logger and exists an application with status code 1.
func Fatal(logger log.Logger, err error, keyval ...interface{}) {
	contextErrorFunc(logger, err).Log(append(keyval, KeyLevel, LevelFatal, KeyMessage, err.Error(), KeyTimestamp, timestampFunc())...)
	os.Exit(1)
}

// Panic log error using given logger and panics.
func Panic(logger log.Logger, err error, keyval ...interface{}) {
	contextErrorFunc(logger, err).Log(append(keyval, KeyLevel, LevelPanic, KeyMessage, err.Error(), KeyTimestamp, timestampFunc())...)
	panic(fmt.Sprint(append(keyval, KeyLevel, LevelPanic, KeyMessage, err)...))
}
