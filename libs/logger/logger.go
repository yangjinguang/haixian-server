package logger

import (
	"log"
	"os"
	"fmt"
	"io"
	"github.com/yangjinguang/wechat-server/libs/config"
)

var logger *log.Logger

func init() {
	logFlag := log.Ldate | log.Ltime | log.Lshortfile
	logFile := config.Conf.Logger.File
	var writer io.Writer
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err == nil {
			writer = io.MultiWriter(os.Stdout, file)
		} else {
			log.Panic("error opening file: %v", err)
			writer = os.Stdout
		}
	} else {
		writer = os.Stdout
	}
	logger = log.New(writer, "", logFlag)
}

func getLogLevel() int {
	logLevel := config.Conf.Logger.Level
	if logLevel <= 0 {
		logLevel = 3
	}
	return logLevel
}

func logPrint(lv string, msg string) {
	logSuffix := config.Conf.Logger.Suffix
	if logSuffix == "" {
		logSuffix = "Wechat-Server"
	}
	logger.SetPrefix(fmt.Sprintf("[%s] %s ", logSuffix, lv))
	logger.Output(3, msg)
}

func Debug(v ...interface{}) {
	if getLogLevel() >= 4 {
		logPrint("DEBUG", fmt.Sprint(v...))
	}
}

func Info(v ...interface{}) {
	if getLogLevel() >= 3 {
		logPrint("INFO", fmt.Sprint(v...))
	}
}

func Warn(v ...interface{}) {
	if getLogLevel() >= 2 {
		logPrint("WARN", fmt.Sprint(v...))
	}
}

func Error(v ...interface{}) {
	if getLogLevel() >= 1 {
		logPrint("ERROR", fmt.Sprint(v...))
	}
}
