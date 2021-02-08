package zap

import (
	"github.com/mattfarina/log-go"
	"go.uber.org/zap"
)

// TODO: Add support for non-sugared logger

// NewSugar creates an instance of Zap that wraps a zap sugared logger. It takes
// a preconfigured zap sugar logger as an argument
func NewSugar(lgr *zap.SugaredLogger) *Zap {
	return &Zap{
		logger:       lgr,
		TraceEnabled: false,
	}
}

type Zap struct {
	logger *zap.SugaredLogger

	// TraceEnabled turns on the trace level. This also requires the zap logging
	// level to be set to debut. Zap does not natively support the trace level.
	// It is bubbled up into the debug messages when trace support is turned on.
	TraceEnabled bool
}

// Trace logs a message at the Trace level
func (l Zap) Trace(msg ...interface{}) {
	if l.TraceEnabled {
		l.logger.Debug(msg...)
	}
}

// Tracef formats a message according to a format specifier and logs the
// message at the Trace level
func (l Zap) Tracef(template string, args ...interface{}) {
	if l.TraceEnabled {
		l.logger.Debugf(template, args...)
	}
}

// Tracew logs a message at the Trace level along with some additional
// context (key-value pairs)
func (l Zap) Tracew(msg string, fields log.Fields) {
	if l.TraceEnabled {
		l.logger.Debugw(msg, fieldToAny(fields)...)
	}
}

// Debug logs a message at the Debug level
func (l Zap) Debug(msg ...interface{}) {
	l.logger.Debug(msg...)
}

// Debugf formats a message according to a format specifier and logs the
// message at the Debug level
func (l Zap) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

// Debugw logs a message at the Debug level along with some additional
// context (key-value pairs)
func (l Zap) Debugw(msg string, fields log.Fields) {
	l.logger.Debugw(msg, fieldToAny(fields)...)
}

// Info logs a message at the Info level
func (l Zap) Info(msg ...interface{}) {
	l.logger.Info(msg...)
}

// Infof formats a message according to a format specifier and logs the
// message at the Info level
func (l Zap) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

// Infow logs a message at the Info level along with some additional
// context (key-value pairs)
func (l Zap) Infow(msg string, fields log.Fields) {
	l.logger.Infow(msg, fieldToAny(fields)...)
}

// Warn logs a message at the Warn level
func (l Zap) Warn(msg ...interface{}) {
	l.logger.Warn(msg...)
}

// Warnf formats a message according to a format specifier and logs the
// message at the Warning level
func (l Zap) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

// Warnw logs a message at the Warning level along with some additional
// context (key-value pairs)
func (l Zap) Warnw(msg string, fields log.Fields) {
	l.logger.Warnw(msg, fieldToAny(fields)...)
}

// Error logs a message at the Error level
func (l Zap) Error(msg ...interface{}) {
	l.logger.Error(msg...)
}

// Errorf formats a message according to a format specifier and logs the
// message at the Error level
func (l Zap) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

// Errorw logs a message at the Error level along with some additional
// context (key-value pairs)
func (l Zap) Errorw(msg string, fields log.Fields) {
	l.logger.Errorw(msg, fieldToAny(fields)...)
}

// Panic logs a message at the Panic level and panics
func (l Zap) Panic(msg ...interface{}) {
	l.logger.Panic(msg...)
}

// Panicf formats a message according to a format specifier and logs the
// message at the Panic level and then panics
func (l Zap) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

// Panicw logs a message at the Panic level along with some additional
// context (key-value pairs) and then panics
func (l Zap) Panicw(msg string, fields log.Fields) {
	l.logger.Panicw(msg, fieldToAny(fields)...)
}

// Fatal logs a message at the Fatal level and exists the application
func (l Zap) Fatal(msg ...interface{}) {
	l.logger.Fatal(msg...)
}

// Fatalf formats a message according to a format specifier and logs the
// message at the Fatal level and exits the application
func (l Zap) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

// Fatalw logs a message at the Fatal level along with some additional
// context (key-value pairs) and exits the application
func (l Zap) Fatalw(msg string, fields log.Fields) {
	l.logger.Fatalw(msg, fieldToAny(fields)...)
}

func fieldToAny(flds log.Fields) []interface{} {
	var ret []interface{}
	for k, v := range flds {
		ret = append(ret, zap.Any(k, v))
	}

	return ret
}
