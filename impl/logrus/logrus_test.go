package logrus

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mattfarina/log"
	"github.com/sirupsen/logrus"
)

func TestLogrus(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	buf := &bytes.Buffer{}
	logrus.SetOutput(buf)
	lgr := NewLogrus()

	lgr.Debug("test debug")
	if !strings.Contains(buf.String(), `level=debug msg="test debug"`) {
		t.Error("logrus debug not logging correctly")
	}
	buf.Reset()

	lgr.Debugf("Hello %s", "World")
	if !strings.Contains(buf.String(), `level=debug msg="Hello World"`) {
		t.Error("logrus debug not logging correctly")
	}
	buf.Reset()

	lgr.Debugw("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `level=debug msg="foo bar" baz=qux`) {
		t.Log(buf.String())
		t.Error("logrus info not logging correctly")
	}
	buf.Reset()

	lgr.Info("test info")
	if !strings.Contains(buf.String(), `level=info msg="test info"`) {
		t.Error("logrus info not logging correctly")
	}
	buf.Reset()

	lgr.Infof("Hello %s", "World")
	if !strings.Contains(buf.String(), `level=info msg="Hello World"`) {
		t.Error("logrus info not logging correctly")
	}
	buf.Reset()

	lgr.Infow("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `level=info msg="foo bar" baz=qux`) {
		t.Log(buf.String())
		t.Error("logrus info not logging correctly")
	}
	buf.Reset()

	lgr.Warn("test warn")
	if !strings.Contains(buf.String(), `level=warning msg="test warn"`) {
		t.Log(buf.String())
		t.Error("logrus warn not logging correctly")
	}
	buf.Reset()

	lgr.Warnf("Hello %s", "World")
	if !strings.Contains(buf.String(), `level=warning msg="Hello World"`) {
		t.Log(buf.String())
		t.Error("logrus warn not logging correctly")
	}
	buf.Reset()

	lgr.Warnw("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `level=warning msg="foo bar" baz=qux`) {
		t.Log(buf.String())
		t.Error("logrus warn not logging correctly")
	}
	buf.Reset()

	lgr.Error("test error")
	if !strings.Contains(buf.String(), `level=error msg="test error"`) {
		t.Log(buf.String())
		t.Error("logrus error not logging correctly")
	}
	buf.Reset()

	lgr.Errorf("Hello %s", "World")
	if !strings.Contains(buf.String(), `level=error msg="Hello World"`) {
		t.Log(buf.String())
		t.Error("logrus error not logging correctly")
	}
	buf.Reset()

	lgr.Errorw("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `level=error msg="foo bar" baz=qux`) {
		t.Log(buf.String())
		t.Error("logrus error not logging correctly")
	}
	buf.Reset()

	// lgr.Fatal("test fatal")
	// lgr.Fatalf(template string, args ...interface{})
	// lgr.Fatalw(msg string, fields Fields)

}

func TestLogrusInterface(t *testing.T) {
	lgr := NewLogrus()
	testfunc(lgr)
}

func testfunc(l log.Logger) {
	l.Debug("test")
}
