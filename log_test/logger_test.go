package log_test

import (
	"testing"

	"github.com/colefan/sailfish/log"
)

func TestMyLog(t *testing.T) {
	log.Debug("debug", ",haha = ", 1, ",end")
	log.Debugf("debugf: %d", 1)
	log.Error("error")
	log.Errorf("errorf: %s", "v ...interface{}")
	log.Close()
}
