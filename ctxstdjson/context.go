package ctxstdjson

import (
	"encoding/json"

	"github.com/go-kit/kit/log"
	"github.com/piotrkowalczuk/sklog"
)

// NewContextErrorGeneric creates context for given error.
// Performs error type check internally to choose strategy that fits the best.
func NewContextError(logger log.Logger, err error) *log.Context {
	if ctx, ok := err.(sklog.Contexter); ok {
		return log.NewContext(logger).With(ctx.Context())
	}

	switch e := err.(type) {
	case *json.MarshalerError:
		return NewContextJSONMarshalerError(logger, e)
	case *json.InvalidUnmarshalError:
		return NewContextJSONInvalidUnmarshalError(logger, e)
	case *json.UnmarshalFieldError:
		return NewContextJSONUnmarshalFieldError(logger, e)
	case *json.UnmarshalTypeError:
		return NewContextJSONUnmarshalTypeError(logger, e)
	case *json.UnsupportedTypeError:
		return NewContextJSONUnsupportedTypeError(logger, e)
	case *json.UnsupportedValueError:
		return NewContextJSONUnsupportedValueError(logger, e)
	case *json.InvalidUTF8Error:
		return NewContextJSONInvalidUTF8Error(logger, e)
	case *json.SyntaxError:
		return NewContextJSONSyntaxError(logger, e)
	default:
		return sklog.NewContextErrorGeneric(logger, e)
	}
}

// NewContextJSONUnmarshalTypeError ...
func NewContextJSONUnmarshalTypeError(logger log.Logger, e *json.UnmarshalTypeError) *log.Context {
	return sklog.NewContextErrorGeneric(logger, e).With(
		"json_type", e.Type,
		"json_offset", e.Offset,
		"json_value", e.Value,
	)
}

// NewContextJSONMarshalerError ...
func NewContextJSONMarshalerError(logger log.Logger, e *json.MarshalerError) *log.Context {
	return sklog.NewContextErrorGeneric(logger, e).With(
		"json_type", e.Type,
	)
}

// NewContextJSONInvalidUnmarshalError ...
func NewContextJSONInvalidUnmarshalError(logger log.Logger, e *json.InvalidUnmarshalError) *log.Context {
	return sklog.NewContextErrorGeneric(logger, e).With(
		"json_type", e.Type,
	)
}

// NewContextJSONUnsupportedTypeError ...
func NewContextJSONUnsupportedTypeError(logger log.Logger, e *json.UnsupportedTypeError) *log.Context {
	return sklog.NewContextErrorGeneric(logger, e).With(
		"json_type", e.Type,
	)
}

// NewContextJSONUnsupportedValueError ...
func NewContextJSONUnsupportedValueError(logger log.Logger, e *json.UnsupportedValueError) *log.Context {
	return sklog.NewContextErrorGeneric(logger, e).With(
		"json_value", e.Value,
		"json_str", e.Str,
	)
}

// NewContextJSONInvalidUTF8Error ...
func NewContextJSONInvalidUTF8Error(logger log.Logger, e *json.InvalidUTF8Error) *log.Context {
	return sklog.NewContextErrorGeneric(logger, e).With(
		"json_s", e.S,
	)
}

// NewContextJSONSyntaxError ...
func NewContextJSONSyntaxError(logger log.Logger, e *json.SyntaxError) *log.Context {
	return sklog.NewContextErrorGeneric(logger, e).With(
		"json_offset", e.Offset,
	)
}

// NewContextJSONUnmarshalFieldError ...
func NewContextJSONUnmarshalFieldError(logger log.Logger, e *json.UnmarshalFieldError) *log.Context {
	return sklog.NewContextErrorGeneric(logger, e).With(
		"json_key", e.Key,
		"json_type", e.Type,
		"json_field_name", e.Field.Name,
		"json_field_pkg_path", e.Field.PkgPath,
		"json_field_type", e.Field.Type,
		"json_field_tag", e.Field.Tag,
		"json_field_offset", e.Field.Offset,
		"json_field_index", e.Field.Index,
		"json_field_anonymous", e.Field.Anonymous,
	)
}
