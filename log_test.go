package mylog

import (
	"fmt"
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	log := NewLogger("test", LevelDebug)

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

func TestNewLoggerDefault(t *testing.T) {
	log := NewLoggerDefault("test", LevelDebug, nil)

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

	log.Fatalf("aaa %s aaa", "bbb")
	fmt.Println("cccc")
	fmt.Println("aaaaa")
	time.Sleep(time.Second)
}
