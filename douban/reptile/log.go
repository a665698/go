package reptile

import (
	"log"
	"os"
)

var myLog *log.Logger

func init() {
	myLog = log.New(os.Stdout, "", log.LstdFlags)
}

func NoticeLog(err interface{}) {
	myLog.Println(err)
}
