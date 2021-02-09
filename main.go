package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type log struct {
	logrus.Logger
}

func LogHandler(lvl string) log {
	var l logrus.Level
	switch lvl {
	case "Debug":
		l = logrus.DebugLevel
	case "Info":
		l = logrus.InfoLevel
	case "Error":
		l = logrus.ErrorLevel
	case "Panic":
		l = logrus.PanicLevel
	case "Fatal":
		l = logrus.FatalLevel
	case "WarnL":
		l = logrus.WarnLevel
	case  "Trace":
		l = logrus.TraceLevel
	}
	return log{logrus.Logger{
		Out:          os.Stdout,
		Hooks:        make(logrus.LevelHooks),
		Formatter:    &logrus.JSONFormatter{},
		ReportCaller: false,
		Level:        l,
		ExitFunc:     os.Exit,
	}}

}

func (l log) Log() *logrus.Entry {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	return l.WithField("file", filename).WithField("function", fn)

}

func main() {
	l := LogHandler("Debug")

	l.Log().Debug("test")
}
