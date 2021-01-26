package main

import (
	"github.com/mattfarina/log"
	lgrs "github.com/mattfarina/log/impl/logrus"
	"github.com/sirupsen/logrus"
)

func main() {
	var lg = logrus.New()
	lg.Level = logrus.DebugLevel

	log.Current = lgrs.New(lg)

	log.Debug("Debugging!")

	// A basic message
	log.Info("Hello,", "World")

	// Use Go formatting on a warning
	log.Warnf("Foo %s", "bar")

	// An error with context
	log.Errorw("foo, bar", log.Fields{"baz": "qux"})
}
