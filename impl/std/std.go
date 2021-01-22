package std

import (
	"fmt"
	stdlog "log"

	"github.com/mattfarina/log"
)

func New(lgr *stdlog.Logger) log.Logger {
	return &Logger{
		logger: lgr,
	}
}

func NewStandard() log.Logger {
	return &StdLogger{}
}

type Logger struct {
	logger *stdlog.Logger
}

func (l Logger) Debug(msg string) {
	l.logger.Printf("DEBUG:	%s", msg)
}

func (l Logger) Debugf(template string, args ...interface{}) {
	l.logger.Printf("DEBUG:	"+template, args...)
}

func (l Logger) Debugw(msg string, fields log.Fields) {
	l.logger.Printf("DEBUG:	%s %s", msg, handlFields(fields))
}

func (l Logger) Info(msg string) {
	l.logger.Printf("INFO:	%s", msg)
}

func (l Logger) Infof(template string, args ...interface{}) {
	l.logger.Printf("INFO:	"+template, args...)
}

func (l Logger) Infow(msg string, fields log.Fields) {
	l.logger.Printf("INFO:	%s %s", msg, handlFields(fields))
}

func (l Logger) Warn(msg string) {
	l.logger.Printf("WARNING:	%s", msg)
}

func (l Logger) Warnf(template string, args ...interface{}) {
	l.logger.Printf("WARNING:	"+template, args...)
}

func (l Logger) Warnw(msg string, fields log.Fields) {
	l.logger.Printf("WARNING:	%s %s", msg, handlFields(fields))
}

func (l Logger) Error(msg string) {
	l.logger.Printf("ERROR:	%s", msg)
}

func (l Logger) Errorf(template string, args ...interface{}) {
	l.logger.Printf("ERROR:	"+template, args...)
}

func (l Logger) Errorw(msg string, fields log.Fields) {
	l.logger.Printf("ERROR:	%s %s", msg, handlFields(fields))
}

func (l Logger) Fatal(msg string) {
	l.logger.Fatalf("FATAL:	%s", msg)
}

func (l Logger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf("FATAL:	"+template, args...)
}

func (l Logger) Fatalw(msg string, fields log.Fields) {
	l.logger.Fatalf("FATAL:	%s %s", msg, handlFields(fields))
}

type StdLogger struct{}

func (l StdLogger) Debug(msg string) {
	stdlog.Printf("DEBUG:	%s", msg)
}

func (l StdLogger) Debugf(template string, args ...interface{}) {
	stdlog.Printf("DEBUG:	"+template, args...)
}

func (l StdLogger) Debugw(msg string, fields log.Fields) {
	stdlog.Printf("DEBUG:	%s %s", msg, handlFields(fields))
}

func (l StdLogger) Info(msg string) {
	stdlog.Printf("INFO:	%s", msg)
}

func (l StdLogger) Infof(template string, args ...interface{}) {
	stdlog.Printf("INFO:	"+template, args...)
}

func (l StdLogger) Infow(msg string, fields log.Fields) {
	stdlog.Printf("INFO:	%s %s", msg, handlFields(fields))
}

func (l StdLogger) Warn(msg string) {
	stdlog.Printf("WARNING:	%s", msg)
}

func (l StdLogger) Warnf(template string, args ...interface{}) {
	stdlog.Printf("WARNING:	"+template, args...)
}

func (l StdLogger) Warnw(msg string, fields log.Fields) {
	stdlog.Printf("WARNING:	%s %s", msg, handlFields(fields))
}

func (l StdLogger) Error(msg string) {
	stdlog.Printf("ERROR:	%s", msg)
}

func (l StdLogger) Errorf(template string, args ...interface{}) {
	stdlog.Printf("ERROR:	"+template, args...)
}

func (l StdLogger) Errorw(msg string, fields log.Fields) {
	stdlog.Printf("ERROR:	%s %s", msg, handlFields(fields))
}

func (l StdLogger) Fatal(msg string) {
	stdlog.Fatalf("FATAL:	%s", msg)
}

func (l StdLogger) Fatalf(template string, args ...interface{}) {
	stdlog.Fatalf("FATAL:	"+template, args...)
}

func (l StdLogger) Fatalw(msg string, fields log.Fields) {
	stdlog.Fatalf("FATAL:	%s %s", msg, handlFields(fields))
}

func handlFields(flds log.Fields) string {
	var ret string
	for k, v := range flds {
		ret += fmt.Sprintf("%s=%s ", k, v)
	}
	return ret
}
