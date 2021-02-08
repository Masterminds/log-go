package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/mattfarina/log"
)

// TODO: Add colors support
// TODO: Add a mutex to lock writing output

type Cli struct {
	DebugOut io.Writer
	InfoOut  io.Writer
	WarnOut  io.Writer
	ErrorOut io.Writer
	PanicOut io.Writer
	FatalOut io.Writer

	Level int
}

func NewStandard() log.Logger {
	return &Cli{
		DebugOut: os.Stderr,
		InfoOut:  os.Stdout,
		WarnOut:  os.Stderr,
		ErrorOut: os.Stderr,
		PanicOut: os.Stderr,
		FatalOut: os.Stderr,
		Level:    log.InfoLevel,
	}
}

func (l Cli) Debug(msg ...interface{}) {
	if l.Level <= log.DebugLevel {
		out := fmt.Sprint(append([]interface{}{"DEBUG: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.DebugOut, out)
	}
}

func (l Cli) Debugf(template string, args ...interface{}) {
	if l.Level <= log.DebugLevel {
		out := fmt.Sprintf("DEBUG: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.DebugOut, out)
	}
}

func (l Cli) Debugw(msg string, fields log.Fields) {
	if l.Level <= log.DebugLevel {
		out := fmt.Sprint("DEBUG: "+msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.DebugOut, out)
	}
}

func (l Cli) Info(msg ...interface{}) {
	if l.Level <= log.InfoLevel {
		out := fmt.Sprint(msg...)
		out = checkEnding(out)
		fmt.Fprint(l.InfoOut, out)
	}
}

func (l Cli) Infof(template string, args ...interface{}) {
	if l.Level <= log.InfoLevel {
		out := fmt.Sprintf(template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.InfoOut, out)
	}
}

func (l Cli) Infow(msg string, fields log.Fields) {
	if l.Level <= log.InfoLevel {
		out := fmt.Sprint(msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.InfoOut, out)
	}
}

func (l Cli) Warn(msg ...interface{}) {
	if l.Level <= log.WarnLevel {
		out := fmt.Sprint(append([]interface{}{"WARNING: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.WarnOut, out)
	}
}

func (l Cli) Warnf(template string, args ...interface{}) {
	if l.Level <= log.WarnLevel {
		out := fmt.Sprintf("WARNING: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.WarnOut, out)
	}
}

func (l Cli) Warnw(msg string, fields log.Fields) {
	if l.Level <= log.WarnLevel {
		out := fmt.Sprint("WARNING: "+msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.WarnOut, out)
	}
}

func (l Cli) Error(msg ...interface{}) {
	if l.Level <= log.ErrorLevel {
		out := fmt.Sprint(append([]interface{}{"ERROR: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.ErrorOut, out)
	}
}

func (l Cli) Errorf(template string, args ...interface{}) {
	if l.Level <= log.ErrorLevel {
		out := fmt.Sprintf("ERROR: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.ErrorOut, out)
	}
}

func (l Cli) Errorw(msg string, fields log.Fields) {
	if l.Level <= log.ErrorLevel {
		out := fmt.Sprint("ERROR: "+msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.ErrorOut, out)
	}
}

func (l Cli) Panic(msg ...interface{}) {
	if l.Level <= log.PanicLevel {
		out := fmt.Sprint(append([]interface{}{"PANIC: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.PanicOut, out)
		panic(out)
	}
}

func (l Cli) Panicf(template string, args ...interface{}) {
	if l.Level <= log.PanicLevel {
		out := fmt.Sprintf("PANIC: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.PanicOut, out)
		panic(out)
	}
}

func (l Cli) Panicw(msg string, fields log.Fields) {
	if l.Level <= log.PanicLevel {
		out := fmt.Sprint("PANIC: "+msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.PanicOut, out)
		panic(out)
	}
}

func (l Cli) Fatal(msg ...interface{}) {
	if l.Level <= log.FatalLevel {
		out := fmt.Sprint(append([]interface{}{"FATAL: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.FatalOut, out)
		os.Exit(1)
	}
}

func (l Cli) Fatalf(template string, args ...interface{}) {
	if l.Level <= log.FatalLevel {
		out := fmt.Sprintf("FATAL: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.FatalOut, out)
		os.Exit(1)
	}
}

func (l Cli) Fatalw(msg string, fields log.Fields) {
	if l.Level <= log.FatalLevel {
		out := fmt.Sprint("FATAL: "+msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.FatalOut, out)
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

func checkEnding(in string) string {
	if len(in) == 0 || in[len(in)-1] != '\n' {
		return in + "\n"
	}
	return in
}
