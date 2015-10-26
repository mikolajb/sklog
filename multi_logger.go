package sklog

import "github.com/go-kit/kit/log"

type multiLogger struct {
	loggers []log.Logger
}

// NewMultiLogger returns a Logger that internally fires multiple loggers.
func NewMultiLogger(loggers ...log.Logger) log.Logger {
	return &multiLogger{
		loggers: loggers,
	}
}

// Log implements Logger interface. It returns last error that occurred.
func (ml *multiLogger) Log(keyvals ...interface{}) (err error) {
	for _, logger := range ml.loggers {
		err = logger.Log(keyvals...)
	}

	return
}
