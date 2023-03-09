package interfaces

type Logger interface {
	Error(err error, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
}
