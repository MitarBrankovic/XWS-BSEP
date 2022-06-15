package loggers

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var customLogger = logrus.New()
var successLogger = logrus.New()
var infoLogger = logrus.New()
var errorLogger = logrus.New()

func NewCustomLogger() *logrus.Logger {
	customLogger.SetLevel(logrus.InfoLevel)
	customLogger.SetReportCaller(true)
	customLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "../../logs/custom/customLogs.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	customLogger.SetOutput(multiWriter)
	return customLogger
}

func NewSuccessLogger() *logrus.Logger {
	successLogger.SetLevel(logrus.InfoLevel)
	successLogger.SetReportCaller(true)
	successLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "../../logs/success/successLogs.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	successLogger.SetOutput(multiWriter)
	return successLogger
}

func NewInfoLogger() *logrus.Logger {
	infoLogger.SetLevel(logrus.InfoLevel)
	infoLogger.SetReportCaller(true)
	infoLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "../../logs/info/infoLogs.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	infoLogger.SetOutput(multiWriter)
	return infoLogger
}

func NewErrorLogger() *logrus.Logger {
	infoLogger.SetLevel(logrus.InfoLevel)
	infoLogger.SetReportCaller(true)
	infoLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "../../logs/error/erorLogs.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	infoLogger.SetOutput(multiWriter)
	return infoLogger
}
