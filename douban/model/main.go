package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func Main() error {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:@tcp(127.0.0.1)/movie?charset=utf8")
	if err != nil {
		return err
	}
	mapper := core.NewPrefixMapper(core.SnakeMapper{}, "lnn_")
	engine.SetMapper(mapper)
	return engine.Ping()
}
