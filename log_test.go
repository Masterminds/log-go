package log

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestInterface(t *testing.T) {
	buf := &bytes.Buffer{}
	lgr := newTestLogger(buf)
	Current = lgr

	Trace("test trace")
	if !strings.Contains(buf.String(), `level=trace msg="test trace"`) {
		t.Error("current trace not logging correctly")
	}
	buf.Reset()

	Tracef("Hello %s", "World")
	if !strings.Contains(buf.String(), `level=trace msg="Hello World"`) {
		t.Error("current trace not logging correctly")
	}
	buf.Reset()

	Tracew("foo bar", Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `level=trace msg="foo bar" baz=qux`) {
		t.Log(buf.String())
		t.Error("current trace not logging correctly")
	}
	buf.Reset()

	Debug("test debug")
	if !strings.Contains(buf.String(), `level=debug msg="test debug"`) {
		t.Error("current debug not logging correctly")
	}
	buf.Reset()

	Debugf("Hello %s", "World")
	if !strings.Contains(buf.String(), `level=debug msg="Hello World"`) {
		t.Error("current debug not logging correctly")
	}
	buf.Reset()

	Debugw("foo bar", Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `level=debug msg="foo bar" baz=qux`) {
		t.Log(buf.String())
		t.Error("current info not logging correctly")
	}
	buf.Reset()

	Info("test info")
	if !strings.Contains(buf.String(), `level=info msg="test info"`) {
		t.Error("current info not logging correctly")
	}
	buf.Reset()

	Infof("Hello %s", "World")
	if !strings.Contains(buf.String(), `level=info msg="Hello World"`) {
		t.Error("current info not logging correctly")
	}
	buf.Reset()

	Infow("foo bar", Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `level=info msg="foo bar" baz=qux`) {
		t.Log(buf.String())
		t.Error("current info not logging correctly")
	}
	buf.Reset()

	Warn("test warn")
	if !strings.Contains(buf.String(), `level=warning msg="test warn"`) {
		t.Log(buf.String())
		t.Error("current warn not logging correctly")
	}
	buf.Reset()

	Warnf("Hello %s", "World")
	if !strings.Contains(buf.String(), `level=warning msg="Hello World"`) {
		t.Log(buf.String())
		t.Error("current warn not logging correctly")
	}
	buf.Reset()

	Warnw("foo bar", Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `level=warning msg="foo bar" baz=qux`) {
		t.Log(buf.String())
		t.Error("current warn not logging correctly")
	}
	buf.Reset()

	Error("test error")
	if !strings.Contains(buf.String(), `level=error msg="test error"`) {
		t.Log(buf.String())
		t.Error("current error not logging correctly")
	}
	buf.Reset()

	Errorf("Hello %s", "World")
	if !strings.Contains(buf.String(), `level=error msg="Hello World"`) {
		t.Log(buf.String())
		t.Error("current error not logging correctly")
	}
	buf.Reset()

	Errorw("foo bar", Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `level=error msg="foo bar" baz=qux`) {
		t.Log(buf.String())
		t.Error("current error not logging correctly")
	}
	buf.Reset()

	Panic("test panic")
	if !strings.Contains(buf.String(), `level=panic msg="test panic"`) {
		t.Log(buf.String())
		t.Error("current panic not logging correctly")
	}
	buf.Reset()

	Panicf("Hello %s", "World")
	if !strings.Contains(buf.String(), `level=panic msg="Hello World"`) {
		t.Log(buf.String())
		t.Error("current panic not logging correctly")
	}
	buf.Reset()

	Panicw("foo bar", Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `level=panic msg="foo bar" baz=qux`) {
		t.Log(buf.String())
		t.Error("current panic not logging correctly")
	}
	buf.Reset()

	Fatal("test fatal")
	if !strings.Contains(buf.String(), `level=fatal msg="test fatal"`) {
		t.Log(buf.String())
		t.Error("current error not logging correctly")
	}
	buf.Reset()

	Fatalf("Hello %s", "World")
	if !strings.Contains(buf.String(), `level=fatal msg="Hello World"`) {
		t.Log(buf.String())
		t.Error("current error not logging correctly")
	}
	buf.Reset()

	Fatalw("foo bar", Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `level=fatal msg="foo bar" baz=qux`) {
		t.Log(buf.String())
		t.Error("current error not logging correctly")
	}
	buf.Reset()
}

func TestAtoi(t *testing.T) {
	names := []string{"Trace", "Debug", "Info", "Warn", "Warning", "Error", "Panic", "Fatal"}
	levels := []int{TraceLevel, DebugLevel, InfoLevel, WarnLevel, WarnLevel, ErrorLevel, PanicLevel, FatalLevel}
	if len(names) != len(levels) {
		t.Fatal("inconsistent number of elements in test")
	}
	for n := range names {
		name := names[n]
		level := levels[n]
		testLevel, err := Atoi(name)
		if err != nil {
			t.Errorf("name to level translation returned: %v", err)
		}
		if testLevel != level {
			t.Errorf("incorrect loglevel (%d) set by name \"%s\"", level, name)
		}
	}
	_, err := Atoi("fakeName")
	if err == nil {
		t.Error("fake loglevel name failed to return an error")
	}
}

type testLogger struct {
	logger *bytes.Buffer
}

func newTestLogger(lgr *bytes.Buffer) *testLogger {
	return &testLogger{
		logger: lgr,
	}
}

func (l testLogger) Trace(msg ...interface{}) {
	dummy(l.logger, "trace", msg...)
}

func (l testLogger) Tracef(template string, args ...interface{}) {
	dummyf(l.logger, "trace", template, args...)
}

func (l testLogger) Tracew(msg string, fields Fields) {
	dummyw(l.logger, "trace", msg, fields)
}

func (l testLogger) Debug(msg ...interface{}) {
	dummy(l.logger, "debug", msg...)
}

func (l testLogger) Debugf(template string, args ...interface{}) {
	dummyf(l.logger, "debug", template, args...)
}

func (l testLogger) Debugw(msg string, fields Fields) {
	dummyw(l.logger, "debug", msg, fields)
}

func (l testLogger) Info(msg ...interface{}) {
	dummy(l.logger, "info", msg...)
}

func (l testLogger) Infof(template string, args ...interface{}) {
	dummyf(l.logger, "info", template, args...)
}

func (l testLogger) Infow(msg string, fields Fields) {
	dummyw(l.logger, "info", msg, fields)
}

func (l testLogger) Warn(msg ...interface{}) {
	dummy(l.logger, "warning", msg...)
}

func (l testLogger) Warnf(template string, args ...interface{}) {
	dummyf(l.logger, "warning", template, args...)
}

func (l testLogger) Warnw(msg string, fields Fields) {
	dummyw(l.logger, "warning", msg, fields)
}

func (l testLogger) Error(msg ...interface{}) {
	dummy(l.logger, "error", msg...)
}

func (l testLogger) Errorf(template string, args ...interface{}) {
	dummyf(l.logger, "error", template, args...)
}

func (l testLogger) Errorw(msg string, fields Fields) {
	dummyw(l.logger, "error", msg, fields)
}

func (l testLogger) Panic(msg ...interface{}) {
	dummy(l.logger, "panic", msg...)
}

func (l testLogger) Panicf(template string, args ...interface{}) {
	dummyf(l.logger, "panic", template, args...)
}

func (l testLogger) Panicw(msg string, fields Fields) {
	dummyw(l.logger, "panic", msg, fields)
}

func (l testLogger) Fatal(msg ...interface{}) {
	dummy(l.logger, "fatal", msg...)
}

func (l testLogger) Fatalf(template string, args ...interface{}) {
	dummyf(l.logger, "fatal", template, args...)
}

func (l testLogger) Fatalw(msg string, fields Fields) {
	dummyw(l.logger, "fatal", msg, fields)
}

func dummy(buf *bytes.Buffer, level string, msg ...interface{}) {
	str := fmt.Sprintf(`level=%s msg="%s"`, level, fmt.Sprint(msg...))
	buf.WriteString(str)
}

func dummyf(buf *bytes.Buffer, level, template string, args ...interface{}) {
	str := fmt.Sprintf(template, args...)
	str = fmt.Sprintf(`level=%s msg="%s"`, level, str)
	buf.WriteString(str)
}

func dummyw(buf *bytes.Buffer, level string, msg interface{}, fields Fields) {
	var flds string
	for k, v := range fields {
		flds += fmt.Sprintf("%s=%s ", k, v)
	}
	str := fmt.Sprintf(`level=%s msg="%s" %s`, level, msg, flds)
	buf.WriteString(str)
}
