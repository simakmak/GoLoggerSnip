package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Logger interface {
	Debug(m interface{})
	Debugf(m string, v ...interface{})
	Error(m interface{})
	Errorf(m string, v ...interface{})
	Info(m interface{})
	Infof(m string, v ...interface{})
}

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

func (l log) Debug(m interface{}) {

	l.Logger.Debug(m)
}
func (l log) Debugf(m string, v ...interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	l.WithField("file", filename).WithField("function", fn).Debugf(m, v)
}

func (l log) Error(m interface{}) {
	l.Logger.Error(m)
}

func (l log) Errorf(m string, v ...interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	l.WithField("file", filename).WithField("function", fn).Errorf(m, v)
}

func (l log) Info(m interface{}) {
	l.Logger.Info(m)
}

func (l log) Infof(m string, v ...interface{}) {
	l.Infof(m, v)
}



func main() {
	var l Logger = LogHandler("Debug")

	//l.Info("test")

	l.Debugf("log %s", "test", "test2")
	l.Info( "test")
}
