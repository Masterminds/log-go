package logrus

import (
	"github.com/mattfarina/log-go"
	"github.com/sirupsen/logrus"
)

// Logrus is a logger that wraps the logrus logger and has it conform to the
// log.Logger interface
type Logrus struct {
	logger *logrus.Logger
}

// New takes an existing logrus logger and uses that for logging
func New(lgr *logrus.Logger) *Logrus {
	return &Logrus{
		logger: lgr,
	}
}

// NewStandard returns a logger with a logrus standard logger which it
// instantiates
func NewStandard() *Logrus {
	return &Logrus{
		logger: logrus.StandardLogger(),
	}
}

// Trace logs a message at the Trace level
func (l Logrus) Trace(msg ...interface{}) {
	l.logger.Trace(msg...)
}

// Tracef formats a message according to a format specifier and logs the
// message at the Trace level
func (l Logrus) Tracef(template string, args ...interface{}) {
	l.logger.Tracef(template, args...)
}

// Tracew logs a message at the Trace level along with some additional
// context (key-value pairs)
func (l Logrus) Tracew(msg string, fields log.Fields) {
	l.logger.WithFields(logrus.Fields(fields)).Trace(msg)
}

// Debug logs a message at the Debug level
func (l Logrus) Debug(msg ...interface{}) {
	l.logger.Debug(msg...)
}

// Debugf formats a message according to a format specifier and logs the
// message at the Debug level
func (l Logrus) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

// Debugw logs a message at the Debug level along with some additional
// context (key-value pairs)
func (l Logrus) Debugw(msg string, fields log.Fields) {
	l.logger.WithFields(logrus.Fields(fields)).Debug(msg)
}

// Info logs a message at the Info level
func (l Logrus) Info(msg ...interface{}) {
	l.logger.Info(msg...)
}

// Infof formats a message according to a format specifier and logs the
// message at the Info level
func (l Logrus) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

// Infow logs a message at the Info level along with some additional
// context (key-value pairs)
func (l Logrus) Infow(msg string, fields log.Fields) {
	l.logger.WithFields(logrus.Fields(fields)).Info(msg)
}

// Warn logs a message at the Warn level
func (l Logrus) Warn(msg ...interface{}) {
	l.logger.Warn(msg...)
}

// Warnf formats a message according to a format specifier and logs the
// message at the Warning level
func (l Logrus) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

// Warnw logs a message at the Warning level along with some additional
// context (key-value pairs)
func (l Logrus) Warnw(msg string, fields log.Fields) {
	l.logger.WithFields(logrus.Fields(fields)).Warn(msg)
}

// Error logs a message at the Error level
func (l Logrus) Error(msg ...interface{}) {
	l.logger.Error(msg...)
}

// Errorf formats a message according to a format specifier and logs the
// message at the Error level
func (l Logrus) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

// Errorw logs a message at the Error level along with some additional
// context (key-value pairs)
func (l Logrus) Errorw(msg string, fields log.Fields) {
	l.logger.WithFields(logrus.Fields(fields)).Error(msg)
}

// Panic logs a message at the Panic level and panics
func (l Logrus) Panic(msg ...interface{}) {
	l.logger.Panic(msg...)
}

// Panicf formats a message according to a format specifier and logs the
// message at the Panic level and then panics
func (l Logrus) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

// Panicw logs a message at the Panic level along with some additional
// context (key-value pairs) and then panics
func (l Logrus) Panicw(msg string, fields log.Fields) {
	l.logger.WithFields(logrus.Fields(fields)).Panic(msg)
}

// Fatal logs a message at the Fatal level and exists the application
func (l Logrus) Fatal(msg ...interface{}) {
	l.logger.Fatal(msg...)
}

// Fatalf formats a message according to a format specifier and logs the
// message at the Fatal level and exits the application
func (l Logrus) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

// Fatalw logs a message at the Fatal level along with some additional
// context (key-value pairs) and exits the application
func (l Logrus) Fatalw(msg string, fields log.Fields) {
	l.logger.WithFields(logrus.Fields(fields)).Fatal(msg)
}
