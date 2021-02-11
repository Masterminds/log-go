package zap

import (
	"fmt"

	"github.com/mattfarina/log-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates an instance of Zap that wraps a zap unsugared logger. It takes
// a preconfigured zap logger as an argument
func New(lgr *zap.Logger) *Logger {
	return &Logger{
		logger: lgr,
	}
}

// Logger is a wrapper about a Zap logger that implements the log.Logger interface
type Logger struct {
	logger *zap.Logger
}

// Trace logs a message at the Trace level
func (l Logger) Trace(msg ...interface{}) {
	if ce := l.logger.Check(zap.DebugLevel-1, fmt.Sprint(msg...)); ce != nil {
		ce.Write()
	}
}

// Tracef formats a message according to a format specifier and logs the
// message at the Trace level
func (l Logger) Tracef(template string, args ...interface{}) {
	if ce := l.logger.Check(zap.DebugLevel-1, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

// Tracew logs a message at the Trace level along with some additional
// context (key-value pairs)
func (l Logger) Tracew(msg string, fields log.Fields) {
	if ce := l.logger.Check(zap.DebugLevel-1, msg); ce != nil {
		ce.Write(fieldToAny(fields)...)
	}
}

// Debug logs a message at the Debug level
func (l Logger) Debug(msg ...interface{}) {
	if ce := l.logger.Check(zap.DebugLevel, fmt.Sprint(msg...)); ce != nil {
		ce.Write()
	}
}

// Debugf formats a message according to a format specifier and logs the
// message at the Debug level
func (l Logger) Debugf(template string, args ...interface{}) {
	if ce := l.logger.Check(zap.DebugLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

// Debugw logs a message at the Debug level along with some additional
// context (key-value pairs)
func (l Logger) Debugw(msg string, fields log.Fields) {
	if ce := l.logger.Check(zap.DebugLevel, msg); ce != nil {
		ce.Write(fieldToAny(fields)...)
	}
}

// Info logs a message at the Info level
func (l Logger) Info(msg ...interface{}) {
	if ce := l.logger.Check(zap.InfoLevel, fmt.Sprint(msg...)); ce != nil {
		ce.Write()
	}
}

// Infof formats a message according to a format specifier and logs the
// message at the Info level
func (l Logger) Infof(template string, args ...interface{}) {
	if ce := l.logger.Check(zap.InfoLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

// Infow logs a message at the Info level along with some additional
// context (key-value pairs)
func (l Logger) Infow(msg string, fields log.Fields) {
	if ce := l.logger.Check(zap.InfoLevel, msg); ce != nil {
		ce.Write(fieldToAny(fields)...)
	}
}

// Warn logs a message at the Warn level
func (l Logger) Warn(msg ...interface{}) {
	if ce := l.logger.Check(zap.WarnLevel, fmt.Sprint(msg...)); ce != nil {
		ce.Write()
	}
}

// Warnf formats a message according to a format specifier and logs the
// message at the Warning level
func (l Logger) Warnf(template string, args ...interface{}) {
	if ce := l.logger.Check(zap.WarnLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

// Warnw logs a message at the Warning level along with some additional
// context (key-value pairs)
func (l Logger) Warnw(msg string, fields log.Fields) {
	if ce := l.logger.Check(zap.WarnLevel, msg); ce != nil {
		ce.Write(fieldToAny(fields)...)
	}
}

// Error logs a message at the Error level
func (l Logger) Error(msg ...interface{}) {
	if ce := l.logger.Check(zap.ErrorLevel, fmt.Sprint(msg...)); ce != nil {
		ce.Write()
	}
}

// Errorf formats a message according to a format specifier and logs the
// message at the Error level
func (l Logger) Errorf(template string, args ...interface{}) {
	if ce := l.logger.Check(zap.ErrorLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

// Errorw logs a message at the Error level along with some additional
// context (key-value pairs)
func (l Logger) Errorw(msg string, fields log.Fields) {
	if ce := l.logger.Check(zap.ErrorLevel, msg); ce != nil {
		ce.Write(fieldToAny(fields)...)
	}
}

// Panic logs a message at the Panic level and panics
func (l Logger) Panic(msg ...interface{}) {
	if ce := l.logger.Check(zap.PanicLevel, fmt.Sprint(msg...)); ce != nil {
		ce.Write()
	}
}

// Panicf formats a message according to a format specifier and logs the
// message at the Panic level and then panics
func (l Logger) Panicf(template string, args ...interface{}) {
	if ce := l.logger.Check(zap.PanicLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

// Panicw logs a message at the Panic level along with some additional
// context (key-value pairs) and then panics
func (l Logger) Panicw(msg string, fields log.Fields) {
	if ce := l.logger.Check(zap.PanicLevel, msg); ce != nil {
		ce.Write(fieldToAny(fields)...)
	}
}

// Fatal logs a message at the Fatal level and exists the application
func (l Logger) Fatal(msg ...interface{}) {
	if ce := l.logger.Check(zap.FatalLevel, fmt.Sprint(msg...)); ce != nil {
		ce.Write()
	}
}

// Fatalf formats a message according to a format specifier and logs the
// message at the Fatal level and exits the application
func (l Logger) Fatalf(template string, args ...interface{}) {
	if ce := l.logger.Check(zap.FatalLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

// Fatalw logs a message at the Fatal level along with some additional
// context (key-value pairs) and exits the application
func (l Logger) Fatalw(msg string, fields log.Fields) {
	if ce := l.logger.Check(zap.FatalLevel, msg); ce != nil {
		ce.Write(fieldToAny(fields)...)
	}
}

func fieldToAny(flds log.Fields) []zapcore.Field {
	var ret []zapcore.Field
	for k, v := range flds {
		ret = append(ret, zap.Any(k, v))
	}

	return ret
}
