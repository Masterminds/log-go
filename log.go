package log

type Fields map[string]interface{}

// Logger is an interface for Logging
type Logger interface {
	Debug(msg ...interface{})
	Debugf(template string, args ...interface{})
	Debugw(msg string, fields Fields)

	Info(msg ...interface{})
	Infof(template string, args ...interface{})
	Infow(msg string, fields Fields)

	Warn(msg ...interface{})
	Warnf(template string, args ...interface{})
	Warnw(msg string, fields Fields)

	Error(msg ...interface{})
	Errorf(template string, args ...interface{})
	Errorw(msg string, fields Fields)

	Fatal(msg ...interface{})
	Fatalf(template string, args ...interface{})
	Fatalw(msg string, fields Fields)

	Panic(msg ...interface{})
	Panicf(template string, args ...interface{})
	Panicw(msg string, fields Fields)
}

var Current Logger

func init() {
	Current = NewStandard()
}

func Debug(msg ...interface{}) {
	Current.Debug(msg...)
}

func Debugf(template string, args ...interface{}) {
	Current.Debugf(template, args...)
}

func Debugw(msg string, fields Fields) {
	Current.Debugw(msg, fields)
}

func Info(msg ...interface{}) {
	Current.Info(msg...)
}

func Infof(template string, args ...interface{}) {
	Current.Infof(template, args...)
}

func Infow(msg string, fields Fields) {
	Current.Infow(msg, fields)
}

func Warn(msg ...interface{}) {
	Current.Warn(msg...)
}

func Warnf(template string, args ...interface{}) {
	Current.Warnf(template, args...)
}

func Warnw(msg string, fields Fields) {
	Current.Warnw(msg, fields)
}

func Error(msg ...interface{}) {
	Current.Error(msg...)
}

func Errorf(template string, args ...interface{}) {
	Current.Errorf(template, args...)
}

func Errorw(msg string, fields Fields) {
	Current.Errorw(msg, fields)
}

func Fatal(msg ...interface{}) {
	Current.Fatal(msg...)
}

func Fatalf(template string, args ...interface{}) {
	Current.Fatalf(template, args...)
}

func Fatalw(msg string, fields Fields) {
	Current.Fatalw(msg, fields)
}

func Panic(msg ...interface{}) {
	Current.Panic(msg...)
}

func Panicf(template string, args ...interface{}) {
	Current.Panicf(template, args...)
}

func Panicw(msg string, fields Fields) {
	Current.Panicw(msg, fields)
}
