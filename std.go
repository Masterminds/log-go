package log

import (
	"fmt"
	stdlog "log"
)

func NewStandard() *StdLogger {
	return &StdLogger{
		Level: InfoLevel,
	}
}

const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

type StdLogger struct {
	Level int
}

func (l StdLogger) Debug(msg ...interface{}) {
	if l.Level <= DebugLevel {
		stdlog.Print(append([]interface{}{"[DEBUG]   "}, msg...)...)
	}
}

func (l StdLogger) Debugf(template string, args ...interface{}) {
	if l.Level <= DebugLevel {
		stdlog.Printf("[DEBUG]   "+template, args...)
	}
}

func (l StdLogger) Debugw(msg string, fields Fields) {
	if l.Level <= DebugLevel {
		stdlog.Printf("[DEBUG]   %s %s", msg, handlFields(fields))
	}
}

func (l StdLogger) Info(msg ...interface{}) {
	if l.Level <= InfoLevel {
		stdlog.Print(append([]interface{}{"[INFO]    "}, msg...)...)
	}
}

func (l StdLogger) Infof(template string, args ...interface{}) {
	if l.Level <= InfoLevel {
		stdlog.Printf("[INFO]    "+template, args...)
	}
}

func (l StdLogger) Infow(msg string, fields Fields) {
	if l.Level <= InfoLevel {
		stdlog.Printf("[INFO]    %s %s", msg, handlFields(fields))
	}
}

func (l StdLogger) Warn(msg ...interface{}) {
	if l.Level <= WarnLevel {
		stdlog.Print(append([]interface{}{"[WARNING] "}, msg...)...)
	}
}

func (l StdLogger) Warnf(template string, args ...interface{}) {
	if l.Level <= WarnLevel {
		stdlog.Printf("[WARNING] "+template, args...)
	}
}

func (l StdLogger) Warnw(msg string, fields Fields) {
	if l.Level <= WarnLevel {
		stdlog.Printf("[WARNING] %s %s", msg, handlFields(fields))
	}
}

func (l StdLogger) Error(msg ...interface{}) {
	if l.Level <= ErrorLevel {
		stdlog.Print(append([]interface{}{"[ERROR]   "}, msg...)...)
	}
}

func (l StdLogger) Errorf(template string, args ...interface{}) {
	if l.Level <= ErrorLevel {
		stdlog.Printf("[ERROR]   "+template, args...)
	}
}

func (l StdLogger) Errorw(msg string, fields Fields) {
	if l.Level <= ErrorLevel {
		stdlog.Printf("[ERROR]   %s %s", msg, handlFields(fields))
	}
}

func (l StdLogger) Fatal(msg ...interface{}) {
	if l.Level <= FatalLevel {
		stdlog.Fatal(append([]interface{}{"[FATAL]   "}, msg...)...)
	}
}

func (l StdLogger) Fatalf(template string, args ...interface{}) {
	if l.Level <= FatalLevel {
		stdlog.Fatalf("[FATAL]   "+template, args...)
	}
}

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
