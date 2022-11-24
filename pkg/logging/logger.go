package logging

import "github.com/AlmazDefourten/goapp/pkg/logging/loggers"

type LoggerInterface interface {
	Trace(args ...interface{})
	Debug(args ...interface{})
	Print(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

func GetInstanceLogger(nameOfLogDir string, nameOfLog string) LoggerInterface {
	return loggers.GetLoggerLorgus(nameOfLogDir, nameOfLog).Entry
}
