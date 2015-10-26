package ctxmgo

import (
	"github.com/go-kit/kit/log"
	"github.com/piotrkowalczuk/sklog"
	"gopkg.in/mgo.v2"
)

// NewContextQueryError ...
func NewContextQueryError(logger log.Logger, err *mgo.QueryError) *log.Context {
	return sklog.NewContextErrorGeneric(logger, err).With(
		"mgo_query_code", err.Code,
		"mgo_query_assertion", err.Assertion,
	)
}

// NewContextErrorGeneric ...
func NewContextErrorGeneric(logger log.Logger, err error) *log.Context {
	if mgoe, ok := err.(*mgo.QueryError); ok {
		return NewContextQueryError(logger, mgoe)
	}

	return sklog.NewContextErrorGeneric(logger, err)
}
