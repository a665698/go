package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var mLog *log.Logger

func init() {
	mLog = log.New(os.Stdout, "", log.LstdFlags)
}

func Panic(l ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		mLog.Panic(file, fmt.Sprintf("第%d行", line), l)
	} else {
		mLog.Panic(l)
	}
}

func Error(l ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		mLog.Println("[Error]", file, fmt.Sprintf("第%d行", line), l)
	} else {
		mLog.Println("[Error]", l)
	}
}

func Notify(l ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		mLog.Println("[Notice]", file, fmt.Sprintf("第%d行", line), l)
	} else {
		mLog.Println("[Notice]", l)
	}
}
