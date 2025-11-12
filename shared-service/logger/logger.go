package logger

import (
	"log"
	"os"
	"path"
	"runtime"
)

var (
	SuccessLogger *log.Logger
	InfoLogger    *log.Logger
	WarnLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func Init() {
	SuccessLogger = log.New(os.Stdout, "SUCCESS: ", log.Ldate|log.Ltime)
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	WarnLogger = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
}

func Success(v string) {
	_, file, line, _ := runtime.Caller(1)
	fileName := path.Base(file)
	SuccessLogger.Printf("[%s:%d] %s\n", fileName, line, v)
}

func Info(v string) {
	_, file, line, _ := runtime.Caller(1)
	fileName := path.Base(file)
	InfoLogger.Printf("[%s:%d] %s\n", fileName, line, v)
}

func Warn(v string) {
	_, file, line, _ := runtime.Caller(1)
	fileName := path.Base(file)
	WarnLogger.Printf("[%s:%d] %s\n", fileName, line, v)
}

func Error(v string) {
	_, file, line, _ := runtime.Caller(1)
	fileName := path.Base(file)
	ErrorLogger.Printf("[%s:%d] %s\n", fileName, line, v)
}
