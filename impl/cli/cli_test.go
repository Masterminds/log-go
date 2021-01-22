package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mattfarina/log"
)

func TestLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	lgr := &Cli{
		DebugOut: buf,
		InfoOut:  buf,
		WarnOut:  buf,
		ErrorOut: buf,
		FatalOut: buf,
	}

	lgr.Debug("test debug")
	if !strings.Contains(buf.String(), `test debug`) {
		t.Log(buf.String())
		t.Error("cli debug not logging correctly")
	}
	buf.Reset()

	lgr.Debugf("Hello %s", "World")
	if !strings.Contains(buf.String(), `Hello World`) {
		t.Error("cli debug not logging correctly")
	}
	buf.Reset()

	lgr.Debugw("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `foo bar baz=qux`) {
		t.Log(buf.String())
		t.Error("cli info not logging correctly")
	}
	buf.Reset()

	lgr.Info("test info")
	if !strings.Contains(buf.String(), `test info`) {
		t.Error("cli info not logging correctly")
	}
	buf.Reset()

	lgr.Infof("Hello %s", "World")
	if !strings.Contains(buf.String(), `Hello World`) {
		t.Error("cli info not logging correctly")
	}
	buf.Reset()

	lgr.Infow("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `foo bar baz=qux`) {
		t.Log(buf.String())
		t.Error("cli info not logging correctly")
	}
	buf.Reset()

	lgr.Warn("test warn")
	if !strings.Contains(buf.String(), `test warn`) {
		t.Log(buf.String())
		t.Error("cli warn not logging correctly")
	}
	buf.Reset()

	lgr.Warnf("Hello %s", "World")
	if !strings.Contains(buf.String(), `Hello World`) {
		t.Log(buf.String())
		t.Error("cli warn not logging correctly")
	}
	buf.Reset()

	lgr.Warnw("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `foo bar baz=qux`) {
		t.Log(buf.String())
		t.Error("cli warn not logging correctly")
	}
	buf.Reset()

	lgr.Error("test error")
	if !strings.Contains(buf.String(), `test error`) {
		t.Log(buf.String())
		t.Error("cli error not logging correctly")
	}
	buf.Reset()

	lgr.Errorf("Hello %s", "World")
	if !strings.Contains(buf.String(), `Hello World`) {
		t.Log(buf.String())
		t.Error("cli error not logging correctly")
	}
	buf.Reset()

	lgr.Errorw("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `foo bar baz=qux`) {
		t.Log(buf.String())
		t.Error("cli error not logging correctly")
	}
	buf.Reset()

	// lgr.Fatal("test fatal")
	// lgr.Fatalf(template string, args ...interface{})
	// lgr.Fatalw(msg string, fields Fields)

}
