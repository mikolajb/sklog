package ctxpq

import (
	"github.com/go-kit/kit/log"
	"github.com/lib/pq"
	"github.com/piotrkowalczuk/sklog"
)

// NewContextError ...
func NewContextError(logger log.Logger, err *pq.Error) *log.Context {
	pqCodeName := ""
	pqCodeClass := ""

	if err.Code != "" && len(err.Code) == 5 {
		pqCodeName = err.Code.Name()
		pqCodeClass = err.Code.Class().Name()
	}
	return sklog.NewContextErrorGeneric(logger, err).With(
		"pq_code", err.Code,
		"pq_code_name", pqCodeName,
		"pq_code_class", pqCodeClass,
		"pq_details", err.Detail,
		"pq_hint", err.Hint,
		"pq_table", err.Table,
		"pq_constraint", err.Constraint,
		"pq_internal_query", err.InternalQuery,
		"pq_where", err.Where,
	)
}

// NewContextErrorGeneric ...
func NewContextErrorGeneric(logger log.Logger, err error) *log.Context {
	if pqe, ok := err.(*pq.Error); ok {
		return NewContextError(logger, pqe)
	}

	return sklog.NewContextErrorGeneric(logger, err)
}
