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

type testLogger struct {
	logger *bytes.Buffer
}

func newTestLogger(lgr *bytes.Buffer) *testLogger {
	return &testLogger{
		logger: lgr,
	}
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
