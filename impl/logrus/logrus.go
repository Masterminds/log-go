package logrus

import (
	"github.com/mattfarina/log"
	"github.com/sirupsen/logrus"
)

type Logrus struct {
	logger *logrus.Logger
}

func New(lgr *logrus.Logger) *Logrus {
	return &Logrus{
		logger: lgr,
	}
}

func NewStandard() *Logrus {
	return &Logrus{
		logger: logrus.StandardLogger(),
	}
}

func (l Logrus) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l Logrus) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l Logrus) Debugw(msg string, fields log.Fields) {
	l.logger.WithFields(logrus.Fields(fields)).Debug(msg)
}

func (l Logrus) Info(msg string) {
	l.logger.Info(msg)
}

func (l Logrus) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l Logrus) Infow(msg string, fields log.Fields) {
	l.logger.WithFields(logrus.Fields(fields)).Info(msg)
}

func (l Logrus) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l Logrus) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l Logrus) Warnw(msg string, fields log.Fields) {
	l.logger.WithFields(logrus.Fields(fields)).Warn(msg)
}

func (l Logrus) Error(msg string) {
	l.logger.Error(msg)
}

func (l Logrus) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l Logrus) Errorw(msg string, fields log.Fields) {
	l.logger.WithFields(logrus.Fields(fields)).Error(msg)
}

func (l Logrus) Fatal(msg string) {
	l.logger.Fatal(msg)
}

func (l Logrus) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func (l Logrus) Fatalw(msg string, fields log.Fields) {
	l.logger.WithFields(logrus.Fields(fields)).Fatal(msg)
}
