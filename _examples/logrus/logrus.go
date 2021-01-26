package main

import (
	"github.com/mattfarina/log"
	"github.com/mattfarina/log/impl/logrus"
)

func main() {
	log.Current = logrus.NewStandard()

	// A basic message
	log.Info("Hello,", "World")

	// Use Go formatting on a warning
	log.Warnf("Foo %s", "bar")

	// An error with context
	log.Errorw("foo, bar", log.Fields{"baz": "qux"})
}
