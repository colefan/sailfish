package network

import (
	"github.com/colefan/logg"
)

var netlogger *logg.BaseLogger

//SetLogger set log
func SetLogger(log *logg.BaseLogger) {
	netlogger = log
}

func netDebug(fmt string, v ...interface{}) {
	if netlogger != nil {
		netlogger.Debug(fmt, v...)
	}
}

func netInfo(fmt string, v ...interface{}) {
	if netlogger != nil {
		netlogger.Info(fmt, v...)
	}
}

func netError(fmt string, v ...interface{}) {
	if netlogger != nil {
		netlogger.Error(fmt, v...)
	}
}

func netWarn(fmt string, v ...interface{}) {
	if netlogger != nil {
		netlogger.Warn(fmt, v...)
	}
}
