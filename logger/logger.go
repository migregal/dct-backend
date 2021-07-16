package logger

import (
	"os"
	"path"
	"runtime"
	"strings"
)

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func ConfigureLogger(logPath string) (*logrus.Logger, error) {
	var filePath string
	if logPath[len(logPath)-1:] == string(os.PathSeparator) {
		filePath = logPath + "log-%Y-%m-%d-%H-%M-%S"
	} else {
		filePath = logPath + string(os.PathSeparator) + "log-%Y-%m-%d-%H-%M-%S"
	}

	rl, err := rotatelogs.New(filePath)

	if err != nil {
		return nil, err
	}

	logger := logrus.New()
	logger.SetOutput(rl)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02-15-04-05-000",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			fn := s[len(s)-1]
			_, file := path.Split(f.File)
			return fn, file
		},
	})

	return logger, nil
}
