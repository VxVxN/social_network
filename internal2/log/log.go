package log

import (
	"io/ioutil"
	slog "log"
	"os"
)

type Logger struct {
	Trace   *slog.Logger
	Debug   *slog.Logger
	Info    *slog.Logger
	Warning *slog.Logger
	Error   *slog.Logger
	Fatal   *slog.Logger
	file    *os.File
}

var ComLog = Init("common.log", false)

func Init(nameFile string, isTest bool) *Logger {
	var sLogger Logger
	var err error
	baseDir := os.Getenv("BASE_DIR")
	writer := ioutil.Discard

	if !isTest {
		sLogger.file, err = os.OpenFile(baseDir+"/logs/"+nameFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			slog.Fatalln("Failed to open log file")
		}
		writer = sLogger.file
	}

	sLogger.Trace = slog.New(ioutil.Discard,
		"TRACE:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	sLogger.Debug = slog.New(writer,
		"DEBUG:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	sLogger.Info = slog.New(writer,
		"INFO:    ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	sLogger.Warning = slog.New(writer,
		"WARNING: ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	sLogger.Error = slog.New(writer,
		"ERROR:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	sLogger.Fatal = slog.New(writer,
		"FATAL:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)
	return &sLogger
}

func (l *Logger) RunTrace() {
	l.Trace = slog.New(l.file,
		"TRACE:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)
}
