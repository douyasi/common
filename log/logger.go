package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var log *logrus.Logger

func Logger() *logrus.Logger {
	return log
}

// initLog
func InitLog(logfile string) {
	file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log = logrus.New()
	mw := io.MultiWriter(os.Stdout, file)
	log.SetReportCaller(true)
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:false,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyFile:  "@file",
			logrus.FieldKeyFunc:  "@caller",
			logrus.FieldKeyMsg:   "@message",
		},
	}

	if err == nil {
		log.SetOutput(mw)
	} else {
		log.Error(err)
	}
}