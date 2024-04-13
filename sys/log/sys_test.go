package log_test

import (
	"testing"
	"time"

	"github.com/zmloong/lotus/sys/log"
)

func Test_sys(t *testing.T) {
	ls, _ := log.NewSys(log.SetFileName("./log.log"), log.SetDebugMode(true), log.SetLoglevel(0))
	time.Sleep(3 * time.Second)
	for i := 0; i < 3; i++ {
		ls.Infof("Info%s", 123)
		ls.Debugf("Debug%s", "一二三")
		ls.Errorf("Error%s", "**//--++")
		ls.Warnf("Warn%s", "#@￥%%%……&！~")
	}
}
