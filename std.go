package log

import (
	"fmt"
	stdlog "log"
)

// NewStandard sets up a basic logger using the general one provided in the Go
// standard library
func NewStandard() *StdLogger {
	return &StdLogger{
		Level: InfoLevel,
	}
}

const (
	// TraceLevel is the constant to use when setting the Trace level for loggers
	// provided by this library.
	TraceLevel = iota

	// DebugLevel is the constant to use when setting the Debug level for loggers
	// provided by this library.
	DebugLevel

	// InfoLevel is the constant to use when setting the Info level for loggers
	// provided by this library.
	InfoLevel

	// WarnLevel is the constant to use when setting the Warn level for loggers
	// provided by this library.
	WarnLevel

	// ErrorLevel is the constant to use when setting the Error level for loggers
	// provided by this library.
	ErrorLevel

	// PanicLevel is the constant to use when setting the Panic level for loggers
	// provided by this library.
	PanicLevel

	// FatalLevel is the constant to use when setting the Fatal level for loggers
	// provided by this library.
	FatalLevel
)

// StdLogger is a struct that wraps the general logger provided by the Go
// standard library and causes it to meet the log.Logger interface
type StdLogger struct {
	Level int
}

// Trace logs a message at the Trace level
func (l StdLogger) Trace(msg ...interface{}) {
	if l.Level <= TraceLevel {
		stdlog.Print(append([]interface{}{"[TRACE]   "}, msg...)...)
	}
}

// Tracef formats a message according to a format specifier and logs the
// message at the Trace level
func (l StdLogger) Tracef(template string, args ...interface{}) {
	if l.Level <= TraceLevel {
		stdlog.Printf("[TRACE]   "+template, args...)
	}
}

// Tracew logs a message at the Trace level along with some additional
// context (key-value pairs)
func (l StdLogger) Tracew(msg string, fields Fields) {
	if l.Level <= TraceLevel {
		stdlog.Printf("[TRACE]   %s %s", msg, handlFields(fields))
	}
}

// Debug logs a message at the Debug level
func (l StdLogger) Debug(msg ...interface{}) {
	if l.Level <= DebugLevel {
		stdlog.Print(append([]interface{}{"[DEBUG]   "}, msg...)...)
	}
}

// Debugf formats a message according to a format specifier and logs the
// message at the Debug level
func (l StdLogger) Debugf(template string, args ...interface{}) {
	if l.Level <= DebugLevel {
		stdlog.Printf("[DEBUG]   "+template, args...)
	}
}

// Debugw logs a message at the Debug level along with some additional
// context (key-value pairs)
func (l StdLogger) Debugw(msg string, fields Fields) {
	if l.Level <= DebugLevel {
		stdlog.Printf("[DEBUG]   %s %s", msg, handlFields(fields))
	}
}

// Info logs a message at the Info level
func (l StdLogger) Info(msg ...interface{}) {
	if l.Level <= InfoLevel {
		stdlog.Print(append([]interface{}{"[INFO]    "}, msg...)...)
	}
}

// Infof formats a message according to a format specifier and logs the
// message at the Info level
func (l StdLogger) Infof(template string, args ...interface{}) {
	if l.Level <= InfoLevel {
		stdlog.Printf("[INFO]    "+template, args...)
	}
}

// Infow logs a message at the Info level along with some additional
// context (key-value pairs)
func (l StdLogger) Infow(msg string, fields Fields) {
	if l.Level <= InfoLevel {
		stdlog.Printf("[INFO]    %s %s", msg, handlFields(fields))
	}
}

// Warn logs a message at the Warn level
func (l StdLogger) Warn(msg ...interface{}) {
	if l.Level <= WarnLevel {
		stdlog.Print(append([]interface{}{"[WARNING] "}, msg...)...)
	}
}

// Warnf formats a message according to a format specifier and logs the
// message at the Warning level
func (l StdLogger) Warnf(template string, args ...interface{}) {
	if l.Level <= WarnLevel {
		stdlog.Printf("[WARNING] "+template, args...)
	}
}

// Warnw logs a message at the Warning level along with some additional
// context (key-value pairs)
func (l StdLogger) Warnw(msg string, fields Fields) {
	if l.Level <= WarnLevel {
		stdlog.Printf("[WARNING] %s %s", msg, handlFields(fields))
	}
}

// Error logs a message at the Error level
func (l StdLogger) Error(msg ...interface{}) {
	if l.Level <= ErrorLevel {
		stdlog.Print(append([]interface{}{"[ERROR]   "}, msg...)...)
	}
}

// Errorf formats a message according to a format specifier and logs the
// message at the Error level
func (l StdLogger) Errorf(template string, args ...interface{}) {
	if l.Level <= ErrorLevel {
		stdlog.Printf("[ERROR]   "+template, args...)
	}
}

// Errorw logs a message at the Error level along with some additional
// context (key-value pairs)
func (l StdLogger) Errorw(msg string, fields Fields) {
	if l.Level <= ErrorLevel {
		stdlog.Printf("[ERROR]   %s %s", msg, handlFields(fields))
	}
}

// Panic logs a message at the Panic level and panics
func (l StdLogger) Panic(msg ...interface{}) {
	if l.Level <= PanicLevel {
		stdlog.Panic(append([]interface{}{"[PANIC]   "}, msg...)...)
	}
}

// Panicf formats a message according to a format specifier and logs the
// message at the Panic level and then panics
func (l StdLogger) Panicf(template string, args ...interface{}) {
	if l.Level <= PanicLevel {
		stdlog.Panicf("[PANIC]   "+template, args...)
	}
}

// Panicw logs a message at the Panic level along with some additional
// context (key-value pairs) and then panics
func (l StdLogger) Panicw(msg string, fields Fields) {
	if l.Level <= PanicLevel {
		stdlog.Panicf("[PANIC]   %s %s", msg, handlFields(fields))
	}
}

// Fatal logs a message at the Fatal level and exists the application
func (l StdLogger) Fatal(msg ...interface{}) {
	if l.Level <= FatalLevel {
		stdlog.Fatal(append([]interface{}{"[FATAL]   "}, msg...)...)
	}
}

// Fatalf formats a message according to a format specifier and logs the
// message at the Fatal level and exits the application
func (l StdLogger) Fatalf(template string, args ...interface{}) {
	if l.Level <= FatalLevel {
		stdlog.Fatalf("[FATAL]   "+template, args...)
	}
}

// Fatalw logs a message at the Fatal level along with some additional
// context (key-value pairs) and exits the application
func (l StdLogger) Fatalw(msg string, fields Fields) {
	if l.Level <= FatalLevel {
		stdlog.Fatalf("[FATAL]   %s %s", msg, handlFields(fields))
	}
}

func handlFields(flds Fields) string {
	var ret string
	for k, v := range flds {
		ret += fmt.Sprintf("[%s=%s]", k, v)
	}
	return ret
}
