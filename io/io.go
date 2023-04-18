// Package io provides a means of turning a log.Logger into an io.Writer for
// a chosen level. An example of this would be:
//
//	import(
//	    "io"
//	    "github.com/Masterminds/log-go"
//	    logio "github.com/Masterminds/log-go/io"
//	)
//
//	func main() {
//	    w := logio.NewCurrentWriter(log.InfoLevel)
//	    io.WriteString(w, "foo")
//	}
package io

import (
	"github.com/Masterminds/log-go"
)

// CurrentWriter uses the current package level logger for io writing
type CurrentWriter struct {
	Level int
}

// NewCurrentWriter creates a new CurrentWriter. The levels that can be passed
// to it are:
// - log.TraceLevel:
// - log.DebugLevel:
// - log.InfoLevel:
// - log.WarnLevel:
// - log.ErrorLevel:
// - log.PanicLevel:
// - log.FatalLevel:
func NewCurrentWriter(level int) *CurrentWriter {
	return &CurrentWriter{
		Level: level,
	}
}

// Write is the write method from the io.Writer interface in the standard lib
func (l CurrentWriter) Write(p []byte) (n int, err error) {
	switch l.Level {
	case log.TraceLevel:
		log.Trace(string(p))
	case log.DebugLevel:
		log.Debug(string(p))
	case log.InfoLevel:
		log.Info(string(p))
	case log.WarnLevel:
		log.Warn(string(p))
	case log.ErrorLevel:
		log.Error(string(p))
	case log.PanicLevel:
		log.Panic(string(p))
	case log.FatalLevel:
		log.Fatal(string(p))
	default:
		log.Panicf("Invalid logger level selected: %d", l.Level)
	}

	return len(p), nil
}

// Writer uses the configured logger for io writing
type Writer struct {
	Logger log.Logger
	Level  int
}

// NewWriter creates a new Writer. It accepts a logger and a level that
// will be written on the io.Writer interface. The levels you can pass in are:
// - log.TraceLevel:
// - log.DebugLevel:
// - log.InfoLevel:
// - log.WarnLevel:
// - log.ErrorLevel:
// - log.PanicLevel:
// - log.FatalLevel:
func NewWriter(lgr log.Logger, level int) *Writer {
	return &Writer{
		Logger: lgr,
		Level:  level,
	}
}

// Write is the write method from the io.Writer interface in the standard lib
func (l Writer) Write(p []byte) (n int, err error) {
	switch l.Level {
	case log.TraceLevel:
		l.Logger.Trace(string(p))
	case log.DebugLevel:
		l.Logger.Debug(string(p))
	case log.InfoLevel:
		l.Logger.Info(string(p))
	case log.WarnLevel:
		l.Logger.Warn(string(p))
	case log.ErrorLevel:
		l.Logger.Error(string(p))
	case log.PanicLevel:
		l.Logger.Panic(string(p))
	case log.FatalLevel:
		l.Logger.Fatal(string(p))
	default:
		l.Logger.Panicf("Invalid logger level selected: %d", l.Level)
	}

	return len(p), nil
}
