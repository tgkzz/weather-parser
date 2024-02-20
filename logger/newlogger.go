package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	InfoLog *log.Logger
	ErrLog  *log.Logger
}

func NewLogger(filename string) (newLogger Logger, err error) {
	file := fmt.Sprintf("%s.log", filename)
	logfile, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return Logger{}, err
	}

	flags := log.Ldate | log.Ltime | log.Lshortfile

	log.SetOutput(logfile)
	log.SetFlags(flags)

	newLogger = Logger{
		InfoLog: log.New(logfile, "INFO: ", flags),
		ErrLog:  log.New(logfile, "ERROR: ", flags),
	}

	return
}
