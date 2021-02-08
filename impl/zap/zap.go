package zap

import (
	"github.com/mattfarina/log"
	"go.uber.org/zap"
)

// TODO: Add support for non-sugared logger

func NewSugar(lgr *zap.SugaredLogger) *Zap {
	return &Zap{
		logger:       lgr,
		TraceEnabled: false,
	}
}

type Zap struct {
	logger *zap.SugaredLogger

	// TraceEnabled turns on the trace level. This also requires the zap logging
	// level to be set to debut. Zap does not natively support the trace level.
	// It is bubbled up into the debug messages when trace support is turned on.
	TraceEnabled bool
}

func (l Zap) Trace(msg ...interface{}) {
	if l.TraceEnabled {
		l.logger.Debug(msg...)
	}
}

func (l Zap) Tracef(template string, args ...interface{}) {
	if l.TraceEnabled {
		l.logger.Debugf(template, args...)
	}
}

func (l Zap) Tracew(msg string, fields log.Fields) {
	if l.TraceEnabled {
		l.logger.Debugw(msg, fieldToAny(fields)...)
	}
}

func (l Zap) Debug(msg ...interface{}) {
	l.logger.Debug(msg...)
}

func (l Zap) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l Zap) Debugw(msg string, fields log.Fields) {
	l.logger.Debugw(msg, fieldToAny(fields)...)
}

func (l Zap) Info(msg ...interface{}) {
	l.logger.Info(msg...)
}

func (l Zap) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l Zap) Infow(msg string, fields log.Fields) {
	l.logger.Infow(msg, fieldToAny(fields)...)
}

func (l Zap) Warn(msg ...interface{}) {
	l.logger.Warn(msg...)
}

func (l Zap) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l Zap) Warnw(msg string, fields log.Fields) {
	l.logger.Warnw(msg, fieldToAny(fields)...)
}

func (l Zap) Error(msg ...interface{}) {
	l.logger.Error(msg...)
}

func (l Zap) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l Zap) Errorw(msg string, fields log.Fields) {
	l.logger.Errorw(msg, fieldToAny(fields)...)
}

func (l Zap) Panic(msg ...interface{}) {
	l.logger.Panic(msg...)
}

func (l Zap) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

func (l Zap) Panicw(msg string, fields log.Fields) {
	l.logger.Panicw(msg, fieldToAny(fields)...)
}

func (l Zap) Fatal(msg ...interface{}) {
	l.logger.Fatal(msg...)
}

func (l Zap) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func (l Zap) Fatalw(msg string, fields log.Fields) {
	l.logger.Fatalw(msg, fieldToAny(fields)...)
}

func fieldToAny(flds log.Fields) []interface{} {
	var ret []interface{}
	for k, v := range flds {
		ret = append(ret, zap.Any(k, v))
	}

	return ret
}
