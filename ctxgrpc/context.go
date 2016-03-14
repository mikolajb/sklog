package ctxgrpc

import (
	"github.com/go-kit/kit/log"
	"github.com/piotrkowalczuk/sklog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// NewContextErrorGeneric ...
func NewContextErrorGeneric(logger log.Logger, err error) *log.Context {
	if code := grpc.Code(err); code != codes.Unknown {
		return log.NewContext(logger).With(sklog.KeyMessage, grpc.ErrorDesc(err), "code", code.String())
	}

	return log.NewContext(logger).With(sklog.KeyMessage, grpc.ErrorDesc(err))
}
