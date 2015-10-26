package ctxstd_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/piotrkowalczuk/sklog"
	"github.com/piotrkowalczuk/sklog/ctxstd"
	"github.com/stretchr/testify/assert"
)

func TestNewContextErrorGeneric(t *testing.T) {
	testNewContextErrorGeneric(t, successJSONCtxErrData)
}

func testNewContextErrorGeneric(t *testing.T, testData map[string]ctxErrData) {
	b := bytes.NewBuffer(nil)
	l := log.NewJSONLogger(b)

	sklog.SetContextErrorFunc(ctxstd.NewContextErrorGeneric)

	for name, data := range testData {
		err := ctxstd.NewContextErrorGeneric(l, data.error).Log("key", "value")
		if assert.NoError(t, err) {
			got := b.String()

			for _, key := range data.keys {
				assert.Contains(t, got, key, name+" does not contain key: "+key)
			}

			for _, value := range data.values {
				assert.Contains(t, got, value, name+" does not contain value: "+value)
			}
		}

		b.Reset()
	}
}

type ctxErrData struct {
	keys   []string
	values []string
	error  error
}

var successJSONCtxErrData = map[string]ctxErrData{
	"InvalidUnmarshalError": {
		error: &json.InvalidUnmarshalError{
			Type: reflect.TypeOf("string"),
		},
		keys: []string{
			"json_type",
		},
		values: []string{
			"string",
		},
	},
	"MarshalerError": {
		error: &json.MarshalerError{
			Type: reflect.TypeOf("string"),
			Err:  errors.New("sklog_test: MarshalerError example"),
		},
		keys: []string{
			"json_type",
		},
		values: []string{
			"string",
		},
	},
	"SyntaxError": {
		error: &json.SyntaxError{
			Offset: 1,
		},
		keys: []string{
			"json_offset",
		},
		values: []string{
			"1",
		},
	},
	"UnmarshalFieldError": {
		error: &json.UnmarshalFieldError{
			Key:  "key",
			Type: reflect.TypeOf("string"),
			Field: reflect.StructField{
				Name:      "structFieldName",
				PkgPath:   "structFieldPkgPath",
				Type:      reflect.TypeOf("string"),
				Tag:       reflect.StructTag("structFieldTag"),
				Offset:    666,
				Index:     []int{1},
				Anonymous: true,
			},
		},
		keys: []string{
			"json_key",
			"json_type",
			"json_field_name",
			"json_field_pkg_path",
			"json_field_type",
			"json_field_tag",
			"json_field_offset",
			"json_field_index",
			"json_field_anonymous",
		},
		values: []string{
			"key",
			"string",
			"structFieldName",
			"structFieldPkgPath",
			"string",
			"structFieldTag",
			"666",
			"[1]",
			"true",
		},
	},
	"UnmarshalTypeError": {
		error: &json.UnmarshalTypeError{
			Offset: 1,
			Type:   reflect.TypeOf("string"),
			Value:  "value",
		},
		keys: []string{
			"json_offset",
			"json_type",
			"json_value",
		},
		values: []string{
			"1",
			"string",
			"value",
		},
	},
	"UnsupportedTypeError": {
		error: &json.UnsupportedTypeError{
			Type: reflect.TypeOf("string"),
		},
		keys: []string{
			"json_type",
		},
		values: []string{
			"string",
		},
	},
	"UnsupportedValueError": {
		error: &json.UnsupportedValueError{
			Value: reflect.ValueOf("value"),
			Str:   "str",
		},
		keys: []string{
			"json_value",
			"json_str",
		},
		values: []string{
			"value",
			"str",
		},
	},
	"InvalidUTF8Error": {
		error: &json.InvalidUTF8Error{
			S: "s",
		},
		keys: []string{
			"json_s",
		},
		values: []string{
			"s",
		},
	},
}
