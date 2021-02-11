package io

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/mattfarina/log-go"
)

func TestCurrentIO(t *testing.T) {
	buf := &bytes.Buffer{}
	lgr := newTestLogger(buf)
	def := log.Current
	log.Current = lgr
	defer func() {
		log.Current = def
	}()

	o := NewCurrentWriter(log.TraceLevel)
	io.WriteString(o, "testing")
	if !strings.Contains(buf.String(), `level=trace msg="testing"`) {
		t.Error("current trace not logging correctly")
	}
	buf.Reset()

	o = NewCurrentWriter(log.DebugLevel)
	io.WriteString(o, "testing")
	if !strings.Contains(buf.String(), `level=debug msg="testing"`) {
		t.Error("current debug not logging correctly")
	}
	buf.Reset()

	o = NewCurrentWriter(log.InfoLevel)
	io.WriteString(o, "testing")
	if !strings.Contains(buf.String(), `level=info msg="testing"`) {
		t.Error("current info not logging correctly")
	}
	buf.Reset()

	o = NewCurrentWriter(log.WarnLevel)
	io.WriteString(o, "testing")
	if !strings.Contains(buf.String(), `level=warning msg="testing"`) {
		t.Error("current warn not logging correctly")
	}
	buf.Reset()

	o = NewCurrentWriter(log.ErrorLevel)
	io.WriteString(o, "testing")
	if !strings.Contains(buf.String(), `level=error msg="testing"`) {
		t.Error("current error not logging correctly")
	}
	buf.Reset()

	o = NewCurrentWriter(log.PanicLevel)
	io.WriteString(o, "testing")
	if !strings.Contains(buf.String(), `level=panic msg="testing"`) {
		t.Error("current panic not logging correctly")
	}
	buf.Reset()

	o = NewCurrentWriter(log.FatalLevel)
	io.WriteString(o, "testing")
	if !strings.Contains(buf.String(), `level=fatal msg="testing"`) {
		t.Error("current fatal not logging correctly")
	}
	buf.Reset()

	// Testing a non-existant level
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
		}
	}()
	log.Current = def // Test logger does not panic. Using one that does
	o = NewCurrentWriter(5000)
	io.WriteString(o, "not happening")
	if !strings.Contains(buf.String(), `level=000 msg="testing"`) {
		t.Error("current fatal not logging correctly")
	}
	buf.Reset()
	t.Error("Not happening test failed to panic")
}

func TestIO(t *testing.T) {
	buf := &bytes.Buffer{}
	lgr := newTestLogger(buf)

	o := NewWriter(lgr, log.TraceLevel)
	io.WriteString(o, "testing")
	if !strings.Contains(buf.String(), `level=trace msg="testing"`) {
		t.Error("current trace not logging correctly")
	}
	buf.Reset()

	o = NewWriter(lgr, log.DebugLevel)
	io.WriteString(o, "testing")
	if !strings.Contains(buf.String(), `level=debug msg="testing"`) {
		t.Error("current debug not logging correctly")
	}
	buf.Reset()

	o = NewWriter(lgr, log.InfoLevel)
	io.WriteString(o, "testing")
	if !strings.Contains(buf.String(), `level=info msg="testing"`) {
		t.Error("current info not logging correctly")
	}
	buf.Reset()

	o = NewWriter(lgr, log.WarnLevel)
	io.WriteString(o, "testing")
	if !strings.Contains(buf.String(), `level=warning msg="testing"`) {
		t.Error("current warn not logging correctly")
	}
	buf.Reset()

	o = NewWriter(lgr, log.ErrorLevel)
	io.WriteString(o, "testing")
	if !strings.Contains(buf.String(), `level=error msg="testing"`) {
		t.Error("current error not logging correctly")
	}
	buf.Reset()

	o = NewWriter(lgr, log.PanicLevel)
	io.WriteString(o, "testing")
	if !strings.Contains(buf.String(), `level=panic msg="testing"`) {
		t.Error("current panic not logging correctly")
	}
	buf.Reset()

	o = NewWriter(lgr, log.FatalLevel)
	io.WriteString(o, "testing")
	if !strings.Contains(buf.String(), `level=fatal msg="testing"`) {
		t.Error("current fatal not logging correctly")
	}
	buf.Reset()

	// Testing a non-existant level
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
		}
	}()
	o = NewWriter(log.Current, 5000) // Test logger does not panic. Using one that does
	io.WriteString(o, "not happening")
	if !strings.Contains(buf.String(), `level=000 msg="testing"`) {
		t.Error("current fatal not logging correctly")
	}
	buf.Reset()
	t.Error("Not happening test failed to panic")
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

func (l testLogger) Tracew(msg string, fields log.Fields) {
	dummyw(l.logger, "trace", msg, fields)
}

func (l testLogger) Debug(msg ...interface{}) {
	dummy(l.logger, "debug", msg...)
}

func (l testLogger) Debugf(template string, args ...interface{}) {
	dummyf(l.logger, "debug", template, args...)
}

func (l testLogger) Debugw(msg string, fields log.Fields) {
	dummyw(l.logger, "debug", msg, fields)
}

func (l testLogger) Info(msg ...interface{}) {
	dummy(l.logger, "info", msg...)
}

func (l testLogger) Infof(template string, args ...interface{}) {
	dummyf(l.logger, "info", template, args...)
}

func (l testLogger) Infow(msg string, fields log.Fields) {
	dummyw(l.logger, "info", msg, fields)
}

func (l testLogger) Warn(msg ...interface{}) {
	dummy(l.logger, "warning", msg...)
}

func (l testLogger) Warnf(template string, args ...interface{}) {
	dummyf(l.logger, "warning", template, args...)
}

func (l testLogger) Warnw(msg string, fields log.Fields) {
	dummyw(l.logger, "warning", msg, fields)
}

func (l testLogger) Error(msg ...interface{}) {
	dummy(l.logger, "error", msg...)
}

func (l testLogger) Errorf(template string, args ...interface{}) {
	dummyf(l.logger, "error", template, args...)
}

func (l testLogger) Errorw(msg string, fields log.Fields) {
	dummyw(l.logger, "error", msg, fields)
}

func (l testLogger) Panic(msg ...interface{}) {
	dummy(l.logger, "panic", msg...)
}

func (l testLogger) Panicf(template string, args ...interface{}) {
	dummyf(l.logger, "panic", template, args...)
}

func (l testLogger) Panicw(msg string, fields log.Fields) {
	dummyw(l.logger, "panic", msg, fields)
}

func (l testLogger) Fatal(msg ...interface{}) {
	dummy(l.logger, "fatal", msg...)
}

func (l testLogger) Fatalf(template string, args ...interface{}) {
	dummyf(l.logger, "fatal", template, args...)
}

func (l testLogger) Fatalw(msg string, fields log.Fields) {
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

func dummyw(buf *bytes.Buffer, level string, msg interface{}, fields log.Fields) {
	var flds string
	for k, v := range fields {
		flds += fmt.Sprintf("%s=%s ", k, v)
	}
	str := fmt.Sprintf(`level=%s msg="%s" %s`, level, msg, flds)
	buf.WriteString(str)
}
