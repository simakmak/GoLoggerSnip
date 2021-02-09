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

func LogHandler(lvl logrus.Level) log {
	return log{logrus.Logger{
		Out:          os.Stdout,
		Hooks:        make(logrus.LevelHooks),
		Formatter:    &logrus.JSONFormatter{},
		ReportCaller: false,
		Level:        lvl,
		ExitFunc:     os.Exit,
	}}

}

func (l log) Debug(m string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	l.WithField("file", filename).WithField("function", fn).Debug(m)

}

func main() {
	l := LogHandler(logrus.InfoLevel)

	l.Debug("test")
}
