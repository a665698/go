package reptile

import (
	"log"
	"os"
)

var myLog *log.Logger

func init() {
	myLog = log.New(os.Stdout, "", log.LstdFlags)
}

func noticeLog(err interface{}) {
	myLog.Println(err)
}
