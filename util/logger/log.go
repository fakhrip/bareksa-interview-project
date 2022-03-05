package logger

import (
	"io"
	"log"
	"os"
)

const (
	INFO int = iota
	WARN
	ERROR
)

type CustomLogger struct {
	LogInfo  *log.Logger
	LogWarn  *log.Logger
	LogError *log.Logger
}

func CreateLogger(file io.Writer) CustomLogger {
	return CustomLogger{
		log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (customLogger *CustomLogger) WriteLog(logType int, messages ...interface{}) {
	for message := range messages {
		switch logType {
		case INFO:
			customLogger.LogInfo.Println(message)
		case WARN:
			customLogger.LogWarn.Println(message)
		case ERROR:
			customLogger.LogError.Println(message)
			os.Exit(1)
		}
	}
}
