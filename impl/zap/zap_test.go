package zap

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mattfarina/log"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestLogrus(t *testing.T) {

	// Test the logger meets the interface
	var _ log.Logger = new(Zap)

	ts := newTestLogSpy(t)
	defer ts.AssertPassed()
	logger := zaptest.NewLogger(ts, zaptest.Level(zap.DebugLevel))
	defer func() {
		err := logger.Sync()
		if err != nil {
			t.Errorf("Error syncing logger: %s", err)
		}
	}()

	sugar := logger.Sugar()
	lgr := NewSugar(sugar)
	lgr.TraceEnabled = true

	lgr.Trace("test trace")
	lgr.Tracef("Hello %s", "World")
	lgr.Tracew("foo bar", log.Fields{"baz": "qux"})
	lgr.Debug("test debug")
	lgr.Debugf("Hello %s", "World")
	lgr.Debugw("foo bar", log.Fields{"baz": "qux"})
	lgr.Info("test info")
	lgr.Infof("Hello %s", "World")
	lgr.Infow("foo bar", log.Fields{"baz": "qux"})
	lgr.Warn("test warn")
	lgr.Warnf("Hello %s", "World")
	lgr.Warnw("foo bar", log.Fields{"baz": "qux"})
	lgr.Error("test error")
	lgr.Errorf("Hello %s", "World")
	lgr.Errorw("foo bar", log.Fields{"baz": "qux"})
	assert.Panics(t, func() {
		lgr.Panic("test panic")
	})
	assert.Panics(t, func() {
		lgr.Panicf("Hello %s", "World")
	})
	assert.Panics(t, func() {
		lgr.Panicw("foo bar", log.Fields{"baz": "qux"})
	})

	ts.AssertMessages(
		"DEBUG	test trace",
		"DEBUG	Hello World",
		"DEBUG	foo bar	{\"baz\": \"qux\"}",
		"DEBUG	test debug",
		"DEBUG	Hello World",
		"DEBUG	foo bar	{\"baz\": \"qux\"}",
		"INFO	test info",
		"INFO	Hello World",
		"INFO	foo bar	{\"baz\": \"qux\"}",
		"WARN	test warn",
		"WARN	Hello World",
		"WARN	foo bar	{\"baz\": \"qux\"}",
		"ERROR	test error",
		"ERROR	Hello World",
		"ERROR	foo bar	{\"baz\": \"qux\"}",
		"PANIC	test panic",
		"PANIC	Hello World",
		"PANIC	foo bar	{\"baz\": \"qux\"}",
	)

	// lgr.Fatal("test fatal")
	// lgr.Fatalf(template string, args ...interface{})
	// lgr.Fatalw(msg string, fields Fields)

}

func TestZapInterface(t *testing.T) {
	logger := zaptest.NewLogger(t, zaptest.Level(zap.DebugLevel))
	defer func() {
		err := logger.Sync()
		if err != nil {
			t.Errorf("Error syncing logger: %s", err)
		}
	}()
	sugar := logger.Sugar()
	lgr := NewSugar(sugar)
	testfunc(lgr)
}

func testfunc(l log.Logger) {
	l.Debug("test")
}

// This last section is taken from zap itself and licensed under the MIT license
type testLogSpy struct {
	testing.TB

	failed   bool
	Messages []string
}

func newTestLogSpy(t testing.TB) *testLogSpy {
	return &testLogSpy{TB: t}
}

func (t *testLogSpy) Fail() {
	t.failed = true
}

func (t *testLogSpy) Failed() bool {
	return t.failed
}

func (t *testLogSpy) FailNow() {
	t.Fail()
	t.TB.FailNow()
}

func (t *testLogSpy) Logf(format string, args ...interface{}) {
	// Log messages are in the format,
	//
	//   2017-10-27T13:03:01.000-0700	DEBUG	your message here	{data here}
	//
	// We strip the first part of these messages because we can't really test
	// for the timestamp from these tests.
	m := fmt.Sprintf(format, args...)
	m = m[strings.IndexByte(m, '\t')+1:]
	t.Messages = append(t.Messages, m)
	t.TB.Log(m)
}

func (t *testLogSpy) AssertMessages(msgs ...string) {
	assert.Equal(t.TB, msgs, t.Messages, "logged messages did not match")
}

func (t *testLogSpy) AssertPassed() {
	t.assertFailed(false, "expected test to pass")
}

func (t *testLogSpy) AssertFailed() {
	t.assertFailed(true, "expected test to fail")
}

func (t *testLogSpy) assertFailed(v bool, msg string) {
	assert.Equal(t.TB, v, t.failed, msg)
}
