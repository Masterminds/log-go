package log

import (
	"bytes"
	stdlog "log"
	"os"
	"strings"
	"testing"
)

func TestStdLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	stdlog.SetOutput(buf)
	defer func() {
		stdlog.SetOutput(os.Stderr)
	}()

	lgr := NewStandard()

	lgr.Debug("test debug")
	if !strings.Contains(buf.String(), `DEBUG:	test debug`) {
		t.Log(buf.String())
		t.Error("stdlib debug not logging correctly")
	}
	buf.Reset()

	lgr.Debugf("Hello %s", "World")
	if !strings.Contains(buf.String(), `DEBUG:	Hello World`) {
		t.Error("stdlib debug not logging correctly")
	}
	buf.Reset()

	lgr.Debugw("foo bar", Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `DEBUG:	foo bar baz=qux`) {
		t.Log(buf.String())
		t.Error("stdlib info not logging correctly")
	}
	buf.Reset()

	lgr.Info("test info")
	if !strings.Contains(buf.String(), `INFO:	test info`) {
		t.Error("stdlib info not logging correctly")
	}
	buf.Reset()

	lgr.Infof("Hello %s", "World")
	if !strings.Contains(buf.String(), `INFO:	Hello World`) {
		t.Error("stdlib info not logging correctly")
	}
	buf.Reset()

	lgr.Infow("foo bar", Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `INFO:	foo bar baz=qux`) {
		t.Log(buf.String())
		t.Error("stdlib info not logging correctly")
	}
	buf.Reset()

	lgr.Warn("test warn")
	if !strings.Contains(buf.String(), `WARNING:	test warn`) {
		t.Log(buf.String())
		t.Error("stdlib warn not logging correctly")
	}
	buf.Reset()

	lgr.Warnf("Hello %s", "World")
	if !strings.Contains(buf.String(), `WARNING:	Hello World`) {
		t.Log(buf.String())
		t.Error("stdlib warn not logging correctly")
	}
	buf.Reset()

	lgr.Warnw("foo bar", Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `WARNING:	foo bar baz=qux`) {
		t.Log(buf.String())
		t.Error("stdlib warn not logging correctly")
	}
	buf.Reset()

	lgr.Error("test error")
	if !strings.Contains(buf.String(), `ERROR:	test error`) {
		t.Log(buf.String())
		t.Error("stdlib error not logging correctly")
	}
	buf.Reset()

	lgr.Errorf("Hello %s", "World")
	if !strings.Contains(buf.String(), `ERROR:	Hello World`) {
		t.Log(buf.String())
		t.Error("stdlib error not logging correctly")
	}
	buf.Reset()

	lgr.Errorw("foo bar", Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `ERROR:	foo bar baz=qux`) {
		t.Log(buf.String())
		t.Error("stdlib error not logging correctly")
	}
	buf.Reset()

	// lgr.Fatal("test fatal")
	// lgr.Fatalf(template string, args ...interface{})
	// lgr.Fatalw(msg string, fields Fields)

}
