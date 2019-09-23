package log

import (
	"fmt"
	"os"

	"github.com/colefan/logg"
)

type DefaultLogger struct {
	*logg.BaseLogger
}

func NewDefaultLogger(cacheSize int) *DefaultLogger {
	l := &DefaultLogger{
		BaseLogger: logg.NewLogger(cacheSize),
	}
	l.BaseLogger.SetAppender("console", "")
	l.BaseLogger.EnableFuncCallDepath(true)
	l.BaseLogger.SetLogFuncCallDepth(4)
	return l
}

func NewDefaultLoggerFromConfigFile(cacheSize int, iniFile string) *DefaultLogger {
	l := &DefaultLogger{
		BaseLogger: logg.NewLogger(cacheSize),
	}
	l.BaseLogger.LoadConfig(iniFile)
	return l
}

func (l *DefaultLogger) Debug(v ...interface{}) {
	l.BaseLogger.Debug(defaultLoggerMsgHandle(v...))
}

func (l *DefaultLogger) Debugf(format string, v ...interface{}) {
	l.BaseLogger.Debug(format, v...)
}

func (l *DefaultLogger) Info(v ...interface{}) {
	l.BaseLogger.Info(defaultLoggerMsgHandle(v...))
}

func (l *DefaultLogger) Infof(format string, v ...interface{}) {
	l.BaseLogger.Info(format, v...)
}

func (l *DefaultLogger) Warn(v ...interface{}) {
	l.BaseLogger.Warn(defaultLoggerMsgHandle(v...))
}

func (l *DefaultLogger) Warnf(format string, v ...interface{}) {
	l.BaseLogger.Warn(format, v...)
}

func (l *DefaultLogger) Error(v ...interface{}) {
	l.BaseLogger.Error(defaultLoggerMsgHandle(v...))
}

func (l *DefaultLogger) Errorf(format string, v ...interface{}) {
	l.BaseLogger.Error(format, v...)
}

func (l *DefaultLogger) Fatal(v ...interface{}) {
	l.BaseLogger.Fatal(defaultLoggerMsgHandle(v...))
	l.BaseLogger.Close()
	os.Exit(1)
}

func (l *DefaultLogger) Fatalf(format string, v ...interface{}) {
	l.BaseLogger.Fatal(defaultLoggerMsgHandle(v...))
	l.BaseLogger.Close()
	os.Exit(1)
}
func (l *DefaultLogger) Close() {
	l.BaseLogger.Close()
}

func (l *DefaultLogger) SetLogFuncCallDepth(d int) {
	if d >= 0 {
		l.BaseLogger.EnableFuncCallDepath(true)
		l.BaseLogger.SetLogFuncCallDepth(d)
	} else {
		l.BaseLogger.EnableFuncCallDepath(false)
	}
}

func defaultLoggerMsgHandle(v ...interface{}) string {
	return fmt.Sprint(v...)
}
