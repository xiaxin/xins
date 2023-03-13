package core

type Logger interface {
	Debug(message string)
	Debugf(template string, args ...interface{})

	Info(message string)
	Infof(template string, args ...interface{})

	Error(message string)
	Errorf(template string, args ...interface{})
}
