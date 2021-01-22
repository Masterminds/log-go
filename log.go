package log

type Fields map[string]interface{}

// Logger is an interface for Logging
type Logger interface {
	Debug(msg string)
	Debugf(template string, args ...interface{})
	Debugw(msg string, fields Fields)

	Info(msg string)
	Infof(template string, args ...interface{})
	Infow(msg string, fields Fields)

	Warn(msg string)
	Warnf(template string, args ...interface{})
	Warnw(msg string, fields Fields)

	Error(msg string)
	Errorf(template string, args ...interface{})
	Errorw(msg string, fields Fields)

	Fatal(msg string)
	Fatalf(template string, args ...interface{})
	Fatalw(msg string, fields Fields)
}

// TODO: Set a default current logger
var Current Logger

func Debug(msg string) {
	Current.Debug(msg)
}

func Debugf(template string, args ...interface{}) {
	Current.Debugf(template, args...)
}

func Debugw(msg string, fields Fields) {
	Current.Debugw(msg, fields)
}

func Info(msg string) {
	Current.Info(msg)
}

func Infof(template string, args ...interface{}) {
	Current.Infof(template, args...)
}

func Infow(msg string, fields Fields) {
	Current.Infow(msg, fields)
}

func Warn(msg string) {
	Current.Warn(msg)
}

func Warnf(template string, args ...interface{}) {
	Current.Warnf(template, args...)
}

func Warnw(msg string, fields Fields) {
	Current.Warnw(msg, fields)
}

func Error(msg string) {
	Current.Error(msg)
}

func Errorf(template string, args ...interface{}) {
	Current.Errorf(template, args...)
}

func Errorw(msg string, fields Fields) {
	Current.Errorw(msg, fields)
}

func Fatal(msg string) {
	Current.Fatal(msg)
}

func Fatalf(template string, args ...interface{}) {
	Current.Fatalf(template, args...)
}

func Fatalw(msg string, fields Fields) {
	Current.Fatalw(msg, fields)
}
