package std

import (
	"fmt"
	stdlog "log"

	"github.com/Masterminds/log-go"
)

// New creates an instance of std.Logger that wraps a logger from the standard
// library. It takes a logger from the standard library as an argument. This
// is now it differs from log.NewStandard() which is used by default. The
// logger is configurable
func New(lgr *stdlog.Logger) *Logger {
	return &Logger{
		logger: lgr,
		Level:  log.InfoLevel,
	}
}

// Logger is a wrapper around an instance of a logger from the Go standard
// library
type Logger struct {
	logger *stdlog.Logger
	Level  int
}

// Trace logs a message at the Trace level
func (l Logger) Trace(msg ...interface{}) {
	if l.Level <= log.TraceLevel {
		l.logger.Print(append([]interface{}{"[TRACE]   "}, msg...)...)
	}
}

// Tracef formats a message according to a format specifier and logs the
// message at the Trace level
func (l Logger) Tracef(template string, args ...interface{}) {
	if l.Level <= log.TraceLevel {
		l.logger.Printf("[TRACE]   "+template, args...)
	}
}

// Tracew logs a message at the Trace level along with some additional
// context (key-value pairs)
func (l Logger) Tracew(msg string, fields log.Fields) {
	if l.Level <= log.TraceLevel {
		l.logger.Printf("[TRACE]   %s %s", msg, handlFields(fields))
	}
}

// Debug logs a message at the Debug level
func (l Logger) Debug(msg ...interface{}) {
	if l.Level <= log.DebugLevel {
		l.logger.Print(append([]interface{}{"[DEBUG]   "}, msg...)...)
	}
}

// Debugf formats a message according to a format specifier and logs the
// message at the Debug level
func (l Logger) Debugf(template string, args ...interface{}) {
	if l.Level <= log.DebugLevel {
		l.logger.Printf("[DEBUG]   "+template, args...)
	}
}

// Debugw logs a message at the Debug level along with some additional
// context (key-value pairs)
func (l Logger) Debugw(msg string, fields log.Fields) {
	if l.Level <= log.DebugLevel {
		l.logger.Printf("[DEBUG]   %s %s", msg, handlFields(fields))
	}
}

// Info logs a message at the Info level
func (l Logger) Info(msg ...interface{}) {
	if l.Level <= log.InfoLevel {
		l.logger.Print(append([]interface{}{"[INFO]    "}, msg...)...)
	}
}

// Infof formats a message according to a format specifier and logs the
// message at the Info level
func (l Logger) Infof(template string, args ...interface{}) {
	if l.Level <= log.InfoLevel {
		l.logger.Printf("[INFO]    "+template, args...)
	}
}

// Infow logs a message at the Info level along with some additional
// context (key-value pairs)
func (l Logger) Infow(msg string, fields log.Fields) {
	if l.Level <= log.InfoLevel {
		l.logger.Printf("[INFO]    %s %s", msg, handlFields(fields))
	}
}

// Warn logs a message at the Warn level
func (l Logger) Warn(msg ...interface{}) {
	if l.Level <= log.WarnLevel {
		l.logger.Print(append([]interface{}{"[WARNING] "}, msg...)...)
	}
}

// Warnf formats a message according to a format specifier and logs the
// message at the Warning level
func (l Logger) Warnf(template string, args ...interface{}) {
	if l.Level <= log.WarnLevel {
		l.logger.Printf("[WARNING] "+template, args...)
	}
}

// Warnw logs a message at the Warning level along with some additional
// context (key-value pairs)
func (l Logger) Warnw(msg string, fields log.Fields) {
	if l.Level <= log.WarnLevel {
		l.logger.Printf("[WARNING] %s %s", msg, handlFields(fields))
	}
}

// Error logs a message at the Error level
func (l Logger) Error(msg ...interface{}) {
	if l.Level <= log.ErrorLevel {
		l.logger.Print(append([]interface{}{"[ERROR]   "}, msg...)...)
	}
}

// Errorf formats a message according to a format specifier and logs the
// message at the Error level
func (l Logger) Errorf(template string, args ...interface{}) {
	if l.Level <= log.ErrorLevel {
		l.logger.Printf("[ERROR]   "+template, args...)
	}
}

// Errorw logs a message at the Error level along with some additional
// context (key-value pairs)
func (l Logger) Errorw(msg string, fields log.Fields) {
	if l.Level <= log.ErrorLevel {
		l.logger.Printf("[ERROR]   %s %s", msg, handlFields(fields))
	}
}

// Panic logs a message at the Panic level and panics
func (l Logger) Panic(msg ...interface{}) {
	if l.Level <= log.PanicLevel {
		l.logger.Panic(append([]interface{}{"[PANIC]   "}, msg...)...)
	}
}

// Panicf formats a message according to a format specifier and logs the
// message at the Panic level and then panics
func (l Logger) Panicf(template string, args ...interface{}) {
	if l.Level <= log.PanicLevel {
		l.logger.Panicf("[PANIC]   "+template, args...)
	}
}

// Panicw logs a message at the Panic level along with some additional
// context (key-value pairs) and then panics
func (l Logger) Panicw(msg string, fields log.Fields) {
	if l.Level <= log.PanicLevel {
		l.logger.Panicf("[PANIC]   %s %s", msg, handlFields(fields))
	}
}

// Fatal logs a message at the Fatal level and exists the application
func (l Logger) Fatal(msg ...interface{}) {
	if l.Level <= log.FatalLevel {
		l.logger.Fatal(append([]interface{}{"[FATAL]   "}, msg...)...)
	}
}

// Fatalf formats a message according to a format specifier and logs the
// message at the Fatal level and exits the application
func (l Logger) Fatalf(template string, args ...interface{}) {
	if l.Level <= log.FatalLevel {
		l.logger.Fatalf("[FATAL]   "+template, args...)
	}
}

// Fatalw logs a message at the Fatal level along with some additional
// context (key-value pairs) and exits the application
func (l Logger) Fatalw(msg string, fields log.Fields) {
	if l.Level <= log.FatalLevel {
		l.logger.Fatalf("[FATAL]   %s %s", msg, handlFields(fields))
	}
}

func handlFields(flds log.Fields) string {
	var ret string
	for k, v := range flds {
		ret += fmt.Sprintf("[%s=%s]", k, v)
	}
	return ret
}
