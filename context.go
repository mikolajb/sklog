package sklog

import "github.com/go-kit/kit/log"

var (
	contextErrorFunc = NewContextErrorGeneric
)

// Contexter is simple wrapper for Context method.
type Contexter interface {
	Context() []interface{}
}

// SetContextErrorFunc sets function that handle unknown type of error by default.
func SetContextErrorFunc(fn func(log.Logger, error) *log.Context) {
	contextErrorFunc = fn
}

// NewContextErrorGeneric allocates context for generic error interface.
func NewContextErrorGeneric(logger log.Logger, err error) *log.Context {
	return log.NewContext(logger).With(KeyMessage, err.Error())
}
