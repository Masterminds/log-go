package log

import (
	"fmt"
	"strings"
)

// Fields represents a map of key-value pairs where the value can be any Go
// type. The value must be able to be converted to a string.
type Fields map[string]interface{}

// Logger is an interface for Logging
type Logger interface {
	// Trace logs a message at the Trace level
	Trace(msg ...interface{})

	// Tracef formats a message according to a format specifier and logs the
	// message at the Trace level
	Tracef(template string, args ...interface{})

	// Tracew logs a message at the Trace level along with some additional
	// context (key-value pairs)
	Tracew(msg string, fields Fields)

	// Debug logs a message at the Debug level
	Debug(msg ...interface{})

	// Debugf formats a message according to a format specifier and logs the
	// message at the Debug level
	Debugf(template string, args ...interface{})

	// Debugw logs a message at the Debug level along with some additional
	// context (key-value pairs)
	Debugw(msg string, fields Fields)

	// Info logs a message at the Info level
	Info(msg ...interface{})

	// Infof formats a message according to a format specifier and logs the
	// message at the Info level
	Infof(template string, args ...interface{})

	// Infow logs a message at the Info level along with some additional
	// context (key-value pairs)
	Infow(msg string, fields Fields)

	// Warn logs a message at the Warn level
	Warn(msg ...interface{})

	// Warnf formats a message according to a format specifier and logs the
	// message at the Warning level
	Warnf(template string, args ...interface{})

	// Warnw logs a message at the Warning level along with some additional
	// context (key-value pairs)
	Warnw(msg string, fields Fields)

	// Error logs a message at the Error level
	Error(msg ...interface{})

	// Errorf formats a message according to a format specifier and logs the
	// message at the Error level
	Errorf(template string, args ...interface{})

	// Errorw logs a message at the Error level along with some additional
	// context (key-value pairs)
	Errorw(msg string, fields Fields)

	// Panic logs a message at the Panic level and panics
	Panic(msg ...interface{})

	// Panicf formats a message according to a format specifier and logs the
	// message at the Panic level and then panics
	Panicf(template string, args ...interface{})

	// Panicw logs a message at the Panic level along with some additional
	// context (key-value pairs) and then panics
	Panicw(msg string, fields Fields)

	// Fatal logs a message at the Fatal level and exists the application
	Fatal(msg ...interface{})

	// Fatalf formats a message according to a format specifier and logs the
	// message at the Fatal level and exits the application
	Fatalf(template string, args ...interface{})

	// Fatalw logs a message at the Fatal level along with some additional
	// context (key-value pairs) and exits the application
	Fatalw(msg string, fields Fields)
}

// Atoi returns the loglevel integer associated with a common loglevel
// string representation.
func Atoi(loglevelStr string) (level int, err error) {
	switch strings.ToLower(loglevelStr) {
	case "trace":
		level = TraceLevel
	case "debug":
		level = DebugLevel
	case "info":
		level = InfoLevel
	case "warn":
		level = WarnLevel
	case "warning":
		level = WarnLevel
	case "error":
		level = ErrorLevel
	case "panic":
		level = PanicLevel
	case "fatal":
		level = FatalLevel
	default:
		err = fmt.Errorf("unknown loglevel: %s", loglevelStr)
	}
	return
}

// Current contains the logger used for the package level logging functions
var Current Logger

func init() {
	Current = NewStandard()
}

// Trace logs a message at the Trace level
func Trace(msg ...interface{}) {
	Current.Trace(msg...)
}

// Tracef formats a message according to a format specifier and logs the
// message at the Trace level
func Tracef(template string, args ...interface{}) {
	Current.Tracef(template, args...)
}

// Tracew logs a message at the Trace level along with some additional
// context (key-value pairs)
func Tracew(msg string, fields Fields) {
	Current.Tracew(msg, fields)
}

// Debug logs a message at the Debug level
func Debug(msg ...interface{}) {
	Current.Debug(msg...)
}

// Debugf formats a message according to a format specifier and logs the
// message at the Debug level
func Debugf(template string, args ...interface{}) {
	Current.Debugf(template, args...)
}

// Debugw logs a message at the Debug level along with some additional
// context (key-value pairs)
func Debugw(msg string, fields Fields) {
	Current.Debugw(msg, fields)
}

// Info logs a message at the Info level
func Info(msg ...interface{}) {
	Current.Info(msg...)
}

// Infof formats a message according to a format specifier and logs the
// message at the Info level
func Infof(template string, args ...interface{}) {
	Current.Infof(template, args...)
}

// Infow logs a message at the Info level along with some additional
// context (key-value pairs)
func Infow(msg string, fields Fields) {
	Current.Infow(msg, fields)
}

// Warn logs a message at the Warn level
func Warn(msg ...interface{}) {
	Current.Warn(msg...)
}

// Warnf formats a message according to a format specifier and logs the
// message at the Warning level
func Warnf(template string, args ...interface{}) {
	Current.Warnf(template, args...)
}

// Warnw logs a message at the Warning level along with some additional
// context (key-value pairs)
func Warnw(msg string, fields Fields) {
	Current.Warnw(msg, fields)
}

// Error logs a message at the Error level
func Error(msg ...interface{}) {
	Current.Error(msg...)
}

// Errorf formats a message according to a format specifier and logs the
// message at the Error level
func Errorf(template string, args ...interface{}) {
	Current.Errorf(template, args...)
}

// Errorw logs a message at the Error level along with some additional
// context (key-value pairs)
func Errorw(msg string, fields Fields) {
	Current.Errorw(msg, fields)
}

// Panic logs a message at the Panic level and panics
func Panic(msg ...interface{}) {
	Current.Panic(msg...)
}

// Panicf formats a message according to a format specifier and logs the
// message at the Panic level and then panics
func Panicf(template string, args ...interface{}) {
	Current.Panicf(template, args...)
}

// Panicw logs a message at the Panic level along with some additional
// context (key-value pairs) and then panics
func Panicw(msg string, fields Fields) {
	Current.Panicw(msg, fields)
}

// Fatal logs a message at the Fatal level and exists the application
func Fatal(msg ...interface{}) {
	Current.Fatal(msg...)
}

// Fatalf formats a message according to a format specifier and logs the
// message at the Fatal level and exits the application
func Fatalf(template string, args ...interface{}) {
	Current.Fatalf(template, args...)
}

// Fatalw logs a message at the Fatal level along with some additional
// context (key-value pairs) and exits the application
func Fatalw(msg string, fields Fields) {
	Current.Fatalw(msg, fields)
}
