package logger

import (
	"buyevent/internal/configs"
	"log"
	"os"
)

// Logger instant for logging operator actions
var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	var (
		infoLoggerFile  *os.File
		errorLoggerFile *os.File
		err             error
	)
	infoLoggerFile, err = os.OpenFile(configs.Configs.InfoLogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	errorLoggerFile, err = os.OpenFile(configs.Configs.ErrorLogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	InfoLogger = log.New(infoLoggerFile, "", log.Ldate|log.Ltime)
	ErrorLogger = log.New(errorLoggerFile, "", log.Ldate|log.Ltime)
}
