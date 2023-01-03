package loggers

import (
	"fmt"
	"github.com/AlmazDefourten/goapp/pkg/logging"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

//writeHook struct for implementation of lorgus hook interface
type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

// Fire implementation of lorgus hook interface
func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

// Levels implementation of lorgus hook interface
func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

type Logger struct {
	*logrus.Entry
}

func GetLoggerLogrus(typeLogger logging.TypeLogger) Logger {
	var entry = Init(typeLogger)
	return Logger{entry}
}

//todo think: its needed?
//func (l *Logger) GetLoggerLorgusWithField(k string, v interface{}) Logger {
//	return Logger{l.WithField(k, v)}
//}

//permission for file
type typePermission uint32

const (
	//todo
	access0777 typePermission = 0777 // -rwxrwxrwx
	access0755 typePermission = 0755 // -rwxr-xr-x
	access0644 typePermission = 0644 // -rw-r–r–
	access0600 typePermission = 0600 // -rw——-
	access0750 typePermission = 0750 // -rwx-r-x—
	access0700 typePermission = 0700 // -rwx——
	access0640 typePermission = 0640 // -rw-r-----
)

const (
	serverLogDir string = "summer"
)

// Init options of logrus
func Init(typeLogger logging.TypeLogger) *logrus.Entry {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", fileName, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}
	pathUtil := logging.LoggerPathUtil{TypeLogger: typeLogger} //utils for getting dir and name of logger's file
	path, err := pathUtil.GetPath()
	if err != nil {
		panic(err.Error())
	}
	err = os.MkdirAll(path.Dir, os.FileMode(access0644))
	if err != nil {
		panic(err)
	}

	allFile, err := os.OpenFile(path.Dir+"/"+path.FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(access0640))
	if err != nil {
		panic(fmt.Sprintf("{"))
	}

	logger.SetOutput(io.Discard)

	logger.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	logger.SetLevel(logrus.TraceLevel)

	return logrus.NewEntry(logger)
}
