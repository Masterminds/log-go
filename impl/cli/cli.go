package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/mattfarina/log-go"
)

// TODO: Add colors support
// TODO: Add a mutex to lock writing output

// Logger provides a CLI based logger. Log messages are written to the CLI as
// terminal style output.
type Logger struct {
	// TraceOut provides a writer to write trace messages to
	TraceOut io.Writer

	// DebugOut provides a writer to write debug messages to
	DebugOut io.Writer

	// InfoOut provides a writer to write info messages to
	InfoOut io.Writer

	// WarnOut provides a writer to write warning messages to
	WarnOut io.Writer

	// ErrorOut provides a writer to write error messages to
	ErrorOut io.Writer

	// PanicOut provides a writer to write panic messages to
	PanicOut io.Writer

	// FatalOut provides a writer to write fatal messages to
	FatalOut io.Writer

	// Level sets the current logging level
	Level int
}

func NewStandard() *Logger {
	return &Logger{
		// Note, stderr is used for all non-info messages by default
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

// Trace logs a message at the Trace level
func (l Logger) Trace(msg ...interface{}) {
	if l.Level <= log.TraceLevel {
		out := fmt.Sprint(append([]interface{}{"TRACE: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.TraceOut, out)
	}
}

// Tracef formats a message according to a format specifier and logs the
// message at the Trace level
func (l Logger) Tracef(template string, args ...interface{}) {
	if l.Level <= log.TraceLevel {
		out := fmt.Sprintf("TRACE: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.TraceOut, out)
	}
}

// Tracew logs a message at the Trace level along with some additional
// context (key-value pairs)
func (l Logger) Tracew(msg string, fields log.Fields) {
	if l.Level <= log.TraceLevel {
		out := fmt.Sprint("TRACE: "+msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.TraceOut, out)
	}
}

// Debug logs a message at the Debug level
func (l Logger) Debug(msg ...interface{}) {
	if l.Level <= log.DebugLevel {
		out := fmt.Sprint(append([]interface{}{"DEBUG: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.DebugOut, out)
	}
}

// Debugf formats a message according to a format specifier and logs the
// message at the Debug level
func (l Logger) Debugf(template string, args ...interface{}) {
	if l.Level <= log.DebugLevel {
		out := fmt.Sprintf("DEBUG: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.DebugOut, out)
	}
}

// Debugw logs a message at the Debug level along with some additional
// context (key-value pairs)
func (l Logger) Debugw(msg string, fields log.Fields) {
	if l.Level <= log.DebugLevel {
		out := fmt.Sprint("DEBUG: "+msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.DebugOut, out)
	}
}

// Info logs a message at the Info level
func (l Logger) Info(msg ...interface{}) {
	if l.Level <= log.InfoLevel {
		out := fmt.Sprint(msg...)
		out = checkEnding(out)
		fmt.Fprint(l.InfoOut, out)
	}
}

// Infof formats a message according to a format specifier and logs the
// message at the Info level
func (l Logger) Infof(template string, args ...interface{}) {
	if l.Level <= log.InfoLevel {
		out := fmt.Sprintf(template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.InfoOut, out)
	}
}

// Infow logs a message at the Info level along with some additional
// context (key-value pairs)
func (l Logger) Infow(msg string, fields log.Fields) {
	if l.Level <= log.InfoLevel {
		out := fmt.Sprint(msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.InfoOut, out)
	}
}

// Warn logs a message at the Warn level
func (l Logger) Warn(msg ...interface{}) {
	if l.Level <= log.WarnLevel {
		out := fmt.Sprint(append([]interface{}{"WARNING: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.WarnOut, out)
	}
}

// Warnf formats a message according to a format specifier and logs the
// message at the Warning level
func (l Logger) Warnf(template string, args ...interface{}) {
	if l.Level <= log.WarnLevel {
		out := fmt.Sprintf("WARNING: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.WarnOut, out)
	}
}

// Warnw logs a message at the Warning level along with some additional
// context (key-value pairs)
func (l Logger) Warnw(msg string, fields log.Fields) {
	if l.Level <= log.WarnLevel {
		out := fmt.Sprint("WARNING: "+msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.WarnOut, out)
	}
}

// Error logs a message at the Error level
func (l Logger) Error(msg ...interface{}) {
	if l.Level <= log.ErrorLevel {
		out := fmt.Sprint(append([]interface{}{"ERROR: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.ErrorOut, out)
	}
}

// Errorf formats a message according to a format specifier and logs the
// message at the Error level
func (l Logger) Errorf(template string, args ...interface{}) {
	if l.Level <= log.ErrorLevel {
		out := fmt.Sprintf("ERROR: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.ErrorOut, out)
	}
}

// Errorw logs a message at the Error level along with some additional
// context (key-value pairs)
func (l Logger) Errorw(msg string, fields log.Fields) {
	if l.Level <= log.ErrorLevel {
		out := fmt.Sprint("ERROR: "+msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.ErrorOut, out)
	}
}

// Panic logs a message at the Panic level and panics
func (l Logger) Panic(msg ...interface{}) {
	if l.Level <= log.PanicLevel {
		out := fmt.Sprint(append([]interface{}{"PANIC: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.PanicOut, out)
		panic(out)
	}
}

// Panicf formats a message according to a format specifier and logs the
// message at the Panic level and then panics
func (l Logger) Panicf(template string, args ...interface{}) {
	if l.Level <= log.PanicLevel {
		out := fmt.Sprintf("PANIC: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.PanicOut, out)
		panic(out)
	}
}

// Panicw logs a message at the Panic level along with some additional
// context (key-value pairs) and then panics
func (l Logger) Panicw(msg string, fields log.Fields) {
	if l.Level <= log.PanicLevel {
		out := fmt.Sprint("PANIC: "+msg, handlFields(fields))
		out = checkEnding(out)
		fmt.Fprint(l.PanicOut, out)
		panic(out)
	}
}

// Fatal logs a message at the Fatal level and exists the application
func (l Logger) Fatal(msg ...interface{}) {
	if l.Level <= log.FatalLevel {
		out := fmt.Sprint(append([]interface{}{"FATAL: "}, msg...)...)
		out = checkEnding(out)
		fmt.Fprint(l.FatalOut, out)
		os.Exit(1)
	}
}

// Fatalf formats a message according to a format specifier and logs the
// message at the Fatal level and exits the application
func (l Logger) Fatalf(template string, args ...interface{}) {
	if l.Level <= log.FatalLevel {
		out := fmt.Sprintf("FATAL: "+template, args...)
		out = checkEnding(out)
		fmt.Fprint(l.FatalOut, out)
		os.Exit(1)
	}
}

// Fatalw logs a message at the Fatal level along with some additional
// context (key-value pairs) and exits the application
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
