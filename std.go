package log

import (
	"fmt"
	stdlog "log"
)

func NewStandard() Logger {
	return &StdLogger{}
}

type StdLogger struct{}

func (l StdLogger) Debug(msg string) {
	stdlog.Printf("DEBUG:	%s", msg)
}

func (l StdLogger) Debugf(template string, args ...interface{}) {
	stdlog.Printf("DEBUG:	"+template, args...)
}

func (l StdLogger) Debugw(msg string, fields Fields) {
	stdlog.Printf("DEBUG:	%s %s", msg, handlFields(fields))
}

func (l StdLogger) Info(msg string) {
	stdlog.Printf("INFO:	%s", msg)
}

func (l StdLogger) Infof(template string, args ...interface{}) {
	stdlog.Printf("INFO:	"+template, args...)
}

func (l StdLogger) Infow(msg string, fields Fields) {
	stdlog.Printf("INFO:	%s %s", msg, handlFields(fields))
}

func (l StdLogger) Warn(msg string) {
	stdlog.Printf("WARNING:	%s", msg)
}

func (l StdLogger) Warnf(template string, args ...interface{}) {
	stdlog.Printf("WARNING:	"+template, args...)
}

func (l StdLogger) Warnw(msg string, fields Fields) {
	stdlog.Printf("WARNING:	%s %s", msg, handlFields(fields))
}

func (l StdLogger) Error(msg string) {
	stdlog.Printf("ERROR:	%s", msg)
}

func (l StdLogger) Errorf(template string, args ...interface{}) {
	stdlog.Printf("ERROR:	"+template, args...)
}

func (l StdLogger) Errorw(msg string, fields Fields) {
	stdlog.Printf("ERROR:	%s %s", msg, handlFields(fields))
}

func (l StdLogger) Fatal(msg string) {
	stdlog.Fatalf("FATAL:	%s", msg)
}

func (l StdLogger) Fatalf(template string, args ...interface{}) {
	stdlog.Fatalf("FATAL:	"+template, args...)
}

func (l StdLogger) Fatalw(msg string, fields Fields) {
	stdlog.Fatalf("FATAL:	%s %s", msg, handlFields(fields))
}

func handlFields(flds Fields) string {
	var ret string
	for k, v := range flds {
		ret += fmt.Sprintf("%s=%s ", k, v)
	}
	return ret
}
