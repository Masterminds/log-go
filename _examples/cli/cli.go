package main

import (
	"github.com/Masterminds/log-go"
	"github.com/Masterminds/log-go/impl/cli"
)

func main() {
	lgr := cli.NewStandard()
	lgr.Level = log.TraceLevel

	log.Current = lgr

	// A basic message
	log.Info("Hello,", "World")

	// A trace message
	log.Trace("A low level trace message")

	// A trace message
	log.Debugf("Hello, %s", "World")

	// Use Go formatting on a warning
	log.Warnf("Foo %s", "bar")

	// An error with context
	log.Errorw("foo, bar", log.Fields{"baz": "qux"})
}
