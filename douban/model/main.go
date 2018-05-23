package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:@tcp(127.0.0.1)/movie?charset=utf8")
	if err != nil {
		panic(err)
	}
	mapper := core.NewPrefixMapper(core.SnakeMapper{}, "lnn_")
	engine.SetTableMapper(mapper)
	engine.ShowSQL()
	err = engine.Ping()
	if err != nil {
		panic(err)
	}
}
