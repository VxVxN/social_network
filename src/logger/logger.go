package logger

import (
	"io/ioutil"
	"log"
	"os"
)

type Logger struct {
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Fatal   *log.Logger
	file    *os.File
}

func Init(nameFile string) *Logger {
	var sLogger Logger
	var err error
	sLogger.file, err = os.OpenFile("/home/vladkmir/go/src/social_network/logs/"+nameFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file")
	}
	sLogger.Trace = log.New(ioutil.Discard,
		"TRACE:   ",
		log.Ldate|log.Ltime|log.Lshortfile)

	sLogger.Info = log.New(sLogger.file,
		"INFO:    ",
		log.Ldate|log.Ltime|log.Lshortfile)

	sLogger.Warning = log.New(sLogger.file,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	sLogger.Error = log.New(sLogger.file,
		"ERROR:   ",
		log.Ldate|log.Ltime|log.Lshortfile)

	sLogger.Fatal = log.New(sLogger.file,
		"FATAL:   ",
		log.Ldate|log.Ltime|log.Lshortfile)
	return &sLogger
}

func (l *Logger) RunTrace() {
	l.Trace = log.New(l.file,
		"TRACE:   ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
