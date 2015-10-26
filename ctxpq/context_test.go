package ctxpq_test

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/lib/pq"
	"github.com/piotrkowalczuk/sklog"
	"github.com/piotrkowalczuk/sklog/ctxpq"
	"github.com/stretchr/testify/assert"
)

var (
	pqErrors = []*pq.Error{
		&pq.Error{Message: "pq_test: example pq errror"},
		&pq.Error{
			Message: "pq_test: example pq errror",
			Code:    pq.ErrorCode("01000"),
			Table:   "example_table",
		},
	}
	genericErrors = []error{
		errors.New("ctxpq_test: example generic error"),
	}
)

func TestNewContextErrorGeneric(t *testing.T) {
	sklog.SetContextErrorFunc(ctxpq.NewContextErrorGeneric)

	b := bytes.NewBuffer(nil)
	l := log.NewJSONLogger(b)

	for _, e := range pqErrors {
		sklog.Error(l, e)

		assertPqError(t, b, e)
		b.Reset()
	}

	for _, e := range genericErrors {
		sklog.Error(l, e)

		assertGenericError(t, b, e)
		b.Reset()
	}
}

func assertPqError(t *testing.T, s fmt.Stringer, e error) {
	assert.Contains(t, s.String(), "pq_code")
	assert.Contains(t, s.String(), "pq_code_name")
	assert.Contains(t, s.String(), "pq_code_class")
	assert.Contains(t, s.String(), "pq_details")
	assert.Contains(t, s.String(), "pq_hint")
	assert.Contains(t, s.String(), "pq_table")
	assert.Contains(t, s.String(), "pq_constraint")
	assert.Contains(t, s.String(), "pq_internal_query")
	assert.Contains(t, s.String(), "pq_where")
	assert.Contains(t, s.String(), sklog.KeyMessage)
	assert.Contains(t, s.String(), e.Error())
}

func assertGenericError(t *testing.T, s fmt.Stringer, e error) {
	assert.NotContains(t, s.String(), "pq_code")
	assert.NotContains(t, s.String(), "pq_code_name")
	assert.NotContains(t, s.String(), "pq_code_class")
	assert.NotContains(t, s.String(), "pq_details")
	assert.NotContains(t, s.String(), "pq_hint")
	assert.NotContains(t, s.String(), "pq_table")
	assert.NotContains(t, s.String(), "pq_constraint")
	assert.NotContains(t, s.String(), "pq_internal_query")
	assert.NotContains(t, s.String(), "pq_where")
	assert.Contains(t, s.String(), sklog.KeyMessage)
	assert.Contains(t, s.String(), e.Error())
}
