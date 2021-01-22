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
