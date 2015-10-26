package ctxstd

import (
	"encoding/json"
	"go/scanner"
	"net"
	"net/textproto"
	"os"
	"reflect"

	"github.com/go-kit/kit/log"
	"github.com/piotrkowalczuk/sklog"
)

// NewContextErrorGeneric creates context for given error.
// Performs error type check internally to choose strategy that fits the best.
func NewContextErrorGeneric(logger log.Logger, err error) *log.Context {
	if ctx, ok := err.(sklog.Contexter); ok {
		return log.NewContext(logger).With(ctx.Context())
	}

	switch e := err.(type) {
	case *reflect.ValueError:
		return NewContextReflectValueError(logger, e)
	// encoding/json
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
	// os
	case *os.PathError:
		return NewContextOSPathError(logger, e)
	case *os.SyscallError:
		return NewContextOSSyscallError(logger, e)
	case *scanner.Error:
		return NewContextScannerError(logger, e)
	case *net.OpError:
		return NewContextNetOpError(logger, e)
	case *textproto.Error:
		return NewContextTextProtoError(logger, e)
	default:
		return sklog.NewContextErrorGeneric(logger, e)
	}
}

// NewContextReflectValueError ...
func NewContextReflectValueError(logger log.Logger, e *reflect.ValueError) *log.Context {
	return sklog.NewContextErrorGeneric(logger, e).With(
		"reflect_kind", e.Kind,
		"reflect_method", e.Method,
	)
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

// NewContextOSPathError ...
func NewContextOSPathError(logger log.Logger, e *os.PathError) *log.Context {
	return sklog.NewContextErrorGeneric(logger, e).With(
		"os_op", e.Op,
		"os_path", e.Path,
	)
}

// NewContextOSSyscallError ...
func NewContextOSSyscallError(logger log.Logger, e *os.SyscallError) *log.Context {
	return sklog.NewContextErrorGeneric(logger, e).With(
		"os_syscall", e.Syscall,
	)
}

// NewContextScannerError ...
func NewContextScannerError(logger log.Logger, e *scanner.Error) *log.Context {
	return sklog.NewContextErrorGeneric(logger, e).With(
		"scanner_pos", e.Pos,
	)
}

// NewContextNetOpError ...
func NewContextNetOpError(logger log.Logger, e *net.OpError) *log.Context {
	return sklog.NewContextErrorGeneric(logger, e).With(
		"net_addr", e.Addr,
		"net_net", e.Net,
		"net_op", e.Op,
		"net_source", e.Source,
	)
}

// NewContextTextProtoError ...
func NewContextTextProtoError(logger log.Logger, e *textproto.Error) *log.Context {
	return sklog.NewContextErrorGeneric(logger, e).With(
		"textproto_code", e.Code,
	)
}
