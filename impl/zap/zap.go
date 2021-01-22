package zap

import (
	"github.com/mattfarina/log"
	"go.uber.org/zap"
)

// TODO: Add support for non-sugared logger

func NewSugar(lgr *zap.SugaredLogger) log.Logger {
	return &Zap{
		logger: lgr,
	}
}

type Zap struct {
	logger *zap.SugaredLogger
}

func (l Zap) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l Zap) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l Zap) Debugw(msg string, fields log.Fields) {
	l.logger.Debugw(msg, fieldToAny(fields)...)
}

func (l Zap) Info(msg string) {
	l.logger.Info(msg)
}

func (l Zap) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l Zap) Infow(msg string, fields log.Fields) {
	l.logger.Infow(msg, fieldToAny(fields)...)
}

func (l Zap) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l Zap) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l Zap) Warnw(msg string, fields log.Fields) {
	l.logger.Warnw(msg, fieldToAny(fields)...)
}

func (l Zap) Error(msg string) {
	l.logger.Error(msg)
}

func (l Zap) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l Zap) Errorw(msg string, fields log.Fields) {
	l.logger.Errorw(msg, fieldToAny(fields)...)
}

func (l Zap) Fatal(msg string) {
	l.logger.Fatal(msg)
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
