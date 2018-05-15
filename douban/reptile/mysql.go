package reptile

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var mySql *sql.DB

func init() {
	var err error
	mySql, err = sql.Open("mysql", "root:@tcp(127.0.0.1)/movie?charset=utf-8")
	if err != nil {
		noticeLog(err)
	}
}
