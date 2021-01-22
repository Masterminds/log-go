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
}

func NewStandard() log.Logger {
	return &Cli{
		DebugOut: os.Stderr,
		InfoOut:  os.Stdout,
		WarnOut:  os.Stderr,
		ErrorOut: os.Stderr,
		FatalOut: os.Stderr,
	}
}

func (l Cli) Debug(msg string) {
	fmt.Fprint(l.DebugOut, msg)
}

func (l Cli) Debugf(template string, args ...interface{}) {
	fmt.Fprintf(l.DebugOut, template, args...)
}

func (l Cli) Debugw(msg string, fields log.Fields) {
	fmt.Fprint(l.DebugOut, msg, handlFields(fields))
}

func (l Cli) Info(msg string) {
	fmt.Fprint(l.InfoOut, msg)
}

func (l Cli) Infof(template string, args ...interface{}) {
	fmt.Fprintf(l.InfoOut, template, args...)
}

func (l Cli) Infow(msg string, fields log.Fields) {
	fmt.Fprint(l.InfoOut, msg, handlFields(fields))
}

func (l Cli) Warn(msg string) {
	fmt.Fprint(l.WarnOut, msg)
}

func (l Cli) Warnf(template string, args ...interface{}) {
	fmt.Fprintf(l.WarnOut, template, args...)
}

func (l Cli) Warnw(msg string, fields log.Fields) {
	fmt.Fprint(l.WarnOut, msg, handlFields(fields))
}

func (l Cli) Error(msg string) {
	fmt.Fprint(l.ErrorOut, msg)
}

func (l Cli) Errorf(template string, args ...interface{}) {
	fmt.Fprintf(l.ErrorOut, template, args...)
}

func (l Cli) Errorw(msg string, fields log.Fields) {
	fmt.Fprint(l.ErrorOut, msg, handlFields(fields))
}

func (l Cli) Fatal(msg string) {
	fmt.Fprint(l.FatalOut, msg)
	os.Exit(1)
}

func (l Cli) Fatalf(template string, args ...interface{}) {
	fmt.Fprintf(l.FatalOut, template, args...)
	os.Exit(1)
}

func (l Cli) Fatalw(msg string, fields log.Fields) {
	fmt.Fprint(l.FatalOut, msg, handlFields(fields))
	os.Exit(1)
}

func handlFields(flds log.Fields) string {
	ret := " "
	for k, v := range flds {
		ret += fmt.Sprintf("%s=%s ", k, v)
	}
	return ret
}
