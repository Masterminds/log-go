package logrus

import (
	"github.com/mattfarina/log"
	"github.com/sirupsen/logrus"
)

type Logrus struct{}

func NewLogrus() *Logrus {
	return &Logrus{}
}

func (l Logrus) Debug(msg string) {
	logrus.Debug(msg)
}

func (l Logrus) Debugf(template string, args ...interface{}) {
	logrus.Debugf(template, args...)
}

func (l Logrus) Debugw(msg string, fields log.Fields) {
	logrus.WithFields(logrus.Fields(fields)).Debug(msg)
}

func (l Logrus) Info(msg string) {
	logrus.Info(msg)
}

func (l Logrus) Infof(template string, args ...interface{}) {
	logrus.Infof(template, args...)
}

func (l Logrus) Infow(msg string, fields log.Fields) {
	logrus.WithFields(logrus.Fields(fields)).Info(msg)
}

func (l Logrus) Warn(msg string) {
	logrus.Warn(msg)
}

func (l Logrus) Warnf(template string, args ...interface{}) {
	logrus.Warnf(template, args...)
}

func (l Logrus) Warnw(msg string, fields log.Fields) {
	logrus.WithFields(logrus.Fields(fields)).Warn(msg)
}

func (l Logrus) Error(msg string) {
	logrus.Error(msg)
}

func (l Logrus) Errorf(template string, args ...interface{}) {
	logrus.Errorf(template, args...)
}

func (l Logrus) Errorw(msg string, fields log.Fields) {
	logrus.WithFields(logrus.Fields(fields)).Error(msg)
}

func (l Logrus) Fatal(msg string) {
	logrus.Fatal(msg)
}

func (l Logrus) Fatalf(template string, args ...interface{}) {
	logrus.Fatalf(template, args...)
}

func (l Logrus) Fatalw(msg string, fields log.Fields) {
	logrus.WithFields(logrus.Fields(fields)).Fatal(msg)
}
