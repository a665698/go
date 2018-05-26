package common

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var myLog *log.Logger

func init() {
	myLog = log.New(os.Stdout, "", log.LstdFlags)
}

func NoticeLog(err ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		myLog.Println(file, fmt.Sprintf("第%d行", line), err)
	} else {
		myLog.Println(err...)
	}
}
