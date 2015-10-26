# Simple Kit Logger &nbsp;[![Build Status](https://travis-ci.org/piotrkowalczuk/sklog.svg?branch=master)](https://travis-ci.org/piotrkowalczuk/sklog)&nbsp;[![godoc reference](https://godoc.org/github.com/piotrkowalczuk/sklog?status.png)][godoc]

aka **sklog** is a wrapper for [go-kit/log](github.com/go-kit/kit/log) package that adds some shorthands, loggers and context packages so its easier to start:

## Quick Start

```go
package main

import (
	"github.com/piotrkowalczuk/sklog"
	"github.com/go-kit/kit/log"
	"io"
)

var (
	writer io.Writer
)

func main() {
	// allocate writer
	
	logger := log.NewJSONLogger(writer)
	
	sklog.Info(logger, "just an info", "key", "val")
	sklog.Debug(logger, "some debug information", "key", "value")
	sklog.Error(logger, errors.New("example: fake error"), "key", "value")
	// sklog.Fatal(logger, errors.New("example: fake error that exists"), "key", "value")
	// sklog.Panic(logger, errors.New("example: fake error that panics"), "key", "value")
}

```

## Shorthands
	
* **[Info](godoc.org/github.com/piotrkowalczuk/sklog/#Info)** - logs message with `level=info`, `msg=msg` and given keyval's.
* **[Debug](godoc.org/github.com/piotrkowalczuk/sklog/#Debug)** - same like [info](godoc.org/github.com/piotrkowalczuk/sklog/#Info) but with debug level.
* **[Error](godoc.org/github.com/piotrkowalczuk/sklog/#Error)** - logs message with `level=error`, `msg=error.Error()` and it tries to create a context from given error using `NewContextErrorGeneric`. It can be changed using `SetContextErrorFunc`.
* **[Fatal](godoc.org/github.com/piotrkowalczuk/sklog/#Fatal)** - same like [error](godoc.org/github.com/piotrkowalczuk/sklog/#Error) but also exits with code 1.
* **[Panic](godoc.org/github.com/piotrkowalczuk/sklog/#Panic)** - same like [error](godoc.org/github.com/piotrkowalczuk/sklog/#Error) but also panics.

## Context Packages
Each package provide logic necessary to get information from `error` objects.

* ctxstd - standard library context builder, in general imports all packages from stdlib. Made for bootstraping.
* ctxpq - [lib/pq](github.com/lib/pq)
* ctxmgo - [gopkg.in/mgo.v2]("gopkg.in/mgo.v2")


## Loggers
### [Humane Logger](godoc.org/github.com/piotrkowalczuk/sklog/#NewHumaneLogger)
Logger that prints easy to read (for humans) output, usefull for development. Inspired by [Sirupsen/logrus](github.com/Sirupsen/logrus). It provides basic implementation that recognize such keys:

* `timestamp`
* `msg`
* `level`
* `subsystem`
* `http_method`
* `http_status`

But also allow to provide custom list of formatters.

#### Output

```bash
[2015-10-25T13:16:09+01:00] [debug] [api-server] - request processed    http_method=get
```

### [Multi Logger](godoc.org/github.com/piotrkowalczuk/sklog/#NewMultiLogger)
Logger that aggregates multiple loggers into one.