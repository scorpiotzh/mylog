package mylog

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	log := NewLogger("test", LevelNotice)

	log.Notice("aaaaa")
	log.Noticef("aaa %s aaa", "bbb")

	log.Debug("aaaaa")
	log.Debugf("aaa %s aaa", "bbb")

	log.Info("aaaaa")
	log.Infof("aaa %s aaa", "bbb")

	log.Warn("aaaaa")
	log.Warnf("aaa %s aaa", "bbb")

	log.Error("aaaaa")
	log.Errorf("aaa %s aaa", "bbb")

	fmt.Println("a")

	//log.Panic("aaaaa")
	//fmt.Println("b")

	//log.Fatalf("aaa %s aaa","bbb")
	//fmt.Println("cccc")
	//fmt.Println("aaaaa")
}
