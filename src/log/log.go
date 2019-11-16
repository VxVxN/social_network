package log

import (
	"io/ioutil"
	slog "log"
	"os"
)

type logger struct {
	Trace   *slog.Logger
	Info    *slog.Logger
	Warning *slog.Logger
	Error   *slog.Logger
	Fatal   *slog.Logger
	file    *os.File
}

var ComLog *logger

func init() {
	ComLog = Init("common.log")
}

func Init(nameFile string) *logger {
	var sLogger logger
	var err error
	sLogger.file, err = os.OpenFile("logs/"+nameFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Fatalln("Failed to open log file")
	}
	sLogger.Trace = slog.New(ioutil.Discard,
		"TRACE:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	sLogger.Info = slog.New(sLogger.file,
		"INFO:    ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	sLogger.Warning = slog.New(sLogger.file,
		"WARNING: ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	sLogger.Error = slog.New(sLogger.file,
		"ERROR:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	sLogger.Fatal = slog.New(sLogger.file,
		"FATAL:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)
	return &sLogger
}

func (l *logger) RunTrace() {
	l.Trace = slog.New(l.file,
		"TRACE:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)
}
