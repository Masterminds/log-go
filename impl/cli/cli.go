package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/mattfarina/log"
)

// TODO: Add colors support

type Cli struct {
	DebugOut io.Writer
	InfoOut  io.Writer
	WarnOut  io.Writer
	ErrorOut io.Writer
	FatalOut io.Writer

	Level int
}

func NewStandard() log.Logger {
	return &Cli{
		DebugOut: os.Stderr,
		InfoOut:  os.Stdout,
		WarnOut:  os.Stderr,
		ErrorOut: os.Stderr,
		FatalOut: os.Stderr,
		Level:    log.InfoLevel,
	}
}

func (l Cli) Debug(msg ...interface{}) {
	if l.Level <= log.DebugLevel {
		fmt.Fprint(l.DebugOut, fmt.Sprint(append([]interface{}{"DEBUG: "}, msg...)...))
	}
}

func (l Cli) Debugf(template string, args ...interface{}) {
	if l.Level <= log.DebugLevel {
		fmt.Fprintf(l.DebugOut, "DEBUG: "+template, args...)
	}
}

func (l Cli) Debugw(msg string, fields log.Fields) {
	if l.Level <= log.DebugLevel {
		fmt.Fprint(l.DebugOut, "DEBUG: "+msg, handlFields(fields))
	}
}

func (l Cli) Info(msg ...interface{}) {
	if l.Level <= log.InfoLevel {
		fmt.Fprint(l.InfoOut, msg...)
	}
}

func (l Cli) Infof(template string, args ...interface{}) {
	if l.Level <= log.InfoLevel {
		fmt.Fprintf(l.InfoOut, template, args...)
	}
}

func (l Cli) Infow(msg string, fields log.Fields) {
	if l.Level <= log.InfoLevel {
		fmt.Fprint(l.InfoOut, msg, handlFields(fields))
	}
}

func (l Cli) Warn(msg ...interface{}) {
	if l.Level <= log.WarnLevel {
		fmt.Fprint(l.WarnOut, fmt.Sprint(append([]interface{}{"WARNING: "}, msg...)...))
	}
}

func (l Cli) Warnf(template string, args ...interface{}) {
	if l.Level <= log.WarnLevel {
		fmt.Fprintf(l.WarnOut, "WARNING: "+template, args...)
	}
}

func (l Cli) Warnw(msg string, fields log.Fields) {
	if l.Level <= log.WarnLevel {
		fmt.Fprint(l.WarnOut, "WARNING: "+msg, handlFields(fields))
	}
}

func (l Cli) Error(msg ...interface{}) {
	if l.Level <= log.ErrorLevel {
		fmt.Fprint(l.ErrorOut, fmt.Sprint(append([]interface{}{"ERROR: "}, msg...)...))
	}
}

func (l Cli) Errorf(template string, args ...interface{}) {
	if l.Level <= log.ErrorLevel {
		fmt.Fprintf(l.ErrorOut, "ERROR: "+template, args...)
	}
}

func (l Cli) Errorw(msg string, fields log.Fields) {
	if l.Level <= log.ErrorLevel {
		fmt.Fprint(l.ErrorOut, "ERROR: "+msg, handlFields(fields))
	}
}

func (l Cli) Fatal(msg ...interface{}) {
	if l.Level <= log.FatalLevel {
		fmt.Fprint(l.FatalOut, fmt.Sprint(append([]interface{}{"FATAL: "}, msg...)...))
		os.Exit(1)
	}
}

func (l Cli) Fatalf(template string, args ...interface{}) {
	if l.Level <= log.FatalLevel {
		fmt.Fprintf(l.FatalOut, "FATAL: "+template, args...)
		os.Exit(1)
	}
}

func (l Cli) Fatalw(msg string, fields log.Fields) {
	if l.Level <= log.FatalLevel {
		fmt.Fprint(l.FatalOut, "FATAL: "+msg, handlFields(fields))
		os.Exit(1)
	}
}

func handlFields(flds log.Fields) string {
	ret := " "
	for k, v := range flds {
		ret += fmt.Sprintf("%s=%s ", k, v)
	}
	return ret
}
