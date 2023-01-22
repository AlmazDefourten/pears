package logging

type Logger interface {
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
type TypeLogger int

const (
	GlobalLogger  TypeLogger = 1
	ServiceLogger            = 2
)

type LoggerPathUtil struct {
	TypeLogger TypeLogger
}

type PathLogger struct {
	Dir      string
	FileName string
}

func (loggerPathHelper *LoggerPathUtil) GetPath() (*PathLogger, error) {
	switch loggerPathHelper.TypeLogger {
	case GlobalLogger:
		return &PathLogger{"logs/global", "global.log"}, nil
	case ServiceLogger:
		return &PathLogger{"logs/service", "service.log"}, nil
	}
	return &PathLogger{}, &ErrorTypeLogger{}
}

type ErrorTypeLogger struct{}

func (*ErrorTypeLogger) Error() string {
	return "don't supported this type of Logger"
}
