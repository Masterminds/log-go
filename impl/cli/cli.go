package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/mattfarina/log"
)

// TODO: Add colors support
// TODO: Add a mutex to lock writing output

type Logger struct {
	TraceOut io.Writer
	DebugOut io.Writer
	InfoOut  io.Writer
	WarnOut  io.Writer
	ErrorOut io.Writer
	PanicOut io.Writer
	FatalOut io.Writer

	Level int
}

func NewStandard() *Logger {
	return &Logger{
		TraceOut: os.Stderr,
		DebugOut: os.Stderr,
		InfoOut:  os.Stdout,
		WarnOut:  os.Stderr,
		ErrorOut: os.Stderr,
		PanicOut: os.Stderr,
		FatalOut: os.Stderr,
		Level:    log.InfoLevel,
	}
}

func (l Logger) Trace(msg ...interface{}) {
	if l.Level <= log.TraceLevel {
		out := fmt.Sprint(append([]interface{}{"TRACE: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.TraceOut, out)
	}
}

func (l Logger) Tracef(template string, args ...interface{}) {
	if l.Level <= log.TraceLevel {
		out := fmt.Sprintf("TRACE: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.TraceOut, out)
	}
}

func (l Logger) Tracew(msg string, fields log.Fields) {
	if l.Level <= log.TraceLevel {
		out := fmt.Sprint("TRACE: "+msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.TraceOut, out)
	}
}

func (l Logger) Debug(msg ...interface{}) {
	if l.Level <= log.DebugLevel {
		out := fmt.Sprint(append([]interface{}{"DEBUG: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.DebugOut, out)
	}
}

func (l Logger) Debugf(template string, args ...interface{}) {
	if l.Level <= log.DebugLevel {
		out := fmt.Sprintf("DEBUG: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.DebugOut, out)
	}
}

func (l Logger) Debugw(msg string, fields log.Fields) {
	if l.Level <= log.DebugLevel {
		out := fmt.Sprint("DEBUG: "+msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.DebugOut, out)
	}
}

func (l Logger) Info(msg ...interface{}) {
	if l.Level <= log.InfoLevel {
		out := fmt.Sprint(msg...)
		out = checkEnding(out)
		fmt.Fprint(l.InfoOut, out)
	}
}

func (l Logger) Infof(template string, args ...interface{}) {
	if l.Level <= log.InfoLevel {
		out := fmt.Sprintf(template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.InfoOut, out)
	}
}

func (l Logger) Infow(msg string, fields log.Fields) {
	if l.Level <= log.InfoLevel {
		out := fmt.Sprint(msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.InfoOut, out)
	}
}

func (l Logger) Warn(msg ...interface{}) {
	if l.Level <= log.WarnLevel {
		out := fmt.Sprint(append([]interface{}{"WARNING: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.WarnOut, out)
	}
}

func (l Logger) Warnf(template string, args ...interface{}) {
	if l.Level <= log.WarnLevel {
		out := fmt.Sprintf("WARNING: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.WarnOut, out)
	}
}

func (l Logger) Warnw(msg string, fields log.Fields) {
	if l.Level <= log.WarnLevel {
		out := fmt.Sprint("WARNING: "+msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.WarnOut, out)
	}
}

func (l Logger) Error(msg ...interface{}) {
	if l.Level <= log.ErrorLevel {
		out := fmt.Sprint(append([]interface{}{"ERROR: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.ErrorOut, out)
	}
}

func (l Logger) Errorf(template string, args ...interface{}) {
	if l.Level <= log.ErrorLevel {
		out := fmt.Sprintf("ERROR: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.ErrorOut, out)
	}
}

func (l Logger) Errorw(msg string, fields log.Fields) {
	if l.Level <= log.ErrorLevel {
		out := fmt.Sprint("ERROR: "+msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.ErrorOut, out)
	}
}

func (l Logger) Panic(msg ...interface{}) {
	if l.Level <= log.PanicLevel {
		out := fmt.Sprint(append([]interface{}{"PANIC: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.PanicOut, out)
		panic(out)
	}
}

func (l Logger) Panicf(template string, args ...interface{}) {
	if l.Level <= log.PanicLevel {
		out := fmt.Sprintf("PANIC: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.PanicOut, out)
		panic(out)
	}
}

func (l Logger) Panicw(msg string, fields log.Fields) {
	if l.Level <= log.PanicLevel {
		out := fmt.Sprint("PANIC: "+msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.PanicOut, out)
		panic(out)
	}
}

func (l Logger) Fatal(msg ...interface{}) {
	if l.Level <= log.FatalLevel {
		out := fmt.Sprint(append([]interface{}{"FATAL: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.FatalOut, out)
		os.Exit(1)
	}
}

func (l Logger) Fatalf(template string, args ...interface{}) {
	if l.Level <= log.FatalLevel {
		out := fmt.Sprintf("FATAL: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.FatalOut, out)
		os.Exit(1)
	}
}

func (l Logger) Fatalw(msg string, fields log.Fields) {
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
