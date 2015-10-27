package sklog

import (
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
)

// GRPCLogger ...
type GRPCLogger interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
}

type gRPCLogger struct {
	log.Logger
}

// NewGRPCLogger ...
func NewGRPCLogger(logger log.Logger) GRPCLogger {
	return &gRPCLogger{
		Logger: logger,
	}
}

func (grl *gRPCLogger) Fatal(args ...interface{}) {
	Fatal(grl.Logger, errors.New(fmt.Sprint(args...)))
}

func (grl *gRPCLogger) Fatalf(format string, args ...interface{}) {
	Fatal(grl.Logger, fmt.Errorf(format, args...))
}

func (grl *gRPCLogger) Fatalln(args ...interface{}) {
	grl.Fatal(args...)
}

func (grl *gRPCLogger) Print(args ...interface{}) {
	var message string
	for i, arg := range args {
		if i != 0 {
			message += ", "
		}
		message += fmt.Sprint(arg)
	}

	Debug(grl.Logger, message)
}

func (grl *gRPCLogger) Printf(format string, args ...interface{}) {
	Debug(grl.Logger, fmt.Sprintf(format, args...))
}

func (grl *gRPCLogger) Println(args ...interface{}) {
	grl.Print(args...)
}
