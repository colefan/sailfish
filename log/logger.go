package log

var l Logger = NewDefaultLogger(256)

//Logger interface
type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})

	Close()
}

// SetLogger set logger
func SetLogger(logger Logger) {
	if l != nil {
		l.Close()
	}
	l = logger
}

// Debug debug
func Debug(v ...interface{}) {
	l.Debug(v...)
}

// Debugf f
func Debugf(format string, v ...interface{}) {
	l.Debugf(format, v...)
}

// Info info
func Info(v ...interface{}) {
	l.Info(v...)
}

//Infof f
func Infof(format string, v ...interface{}) {
	l.Infof(format, v...)
}

// Warn w
func Warn(v ...interface{}) {
	l.Warn(v...)
}

// Warnf f
func Warnf(format string, v ...interface{}) {
	l.Warnf(format, v...)
}

// Error error
func Error(v ...interface{}) {
	l.Error(v...)
}

// Errorf error
func Errorf(format string, v ...interface{}) {
	l.Errorf(format, v...)
}

// Fatal fatal
func Fatal(v ...interface{}) {
	l.Fatal(v...)
}

// Fatalf fatal
func Fatalf(format string, v ...interface{}) {
	l.Fatalf(format, v...)
}

// Close close log
func Close() {
	l.Close()
}
