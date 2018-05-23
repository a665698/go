package model

import (
	"sync"
)

type Type struct {
	Id         int64
	Name       string
	CreateTime int64 `xorm:"created"`
}

func NewType(name string) *Type {
	return &Type{
		Name: name,
	}
}

func (t *Type) Insert() (int64, error) {
	return engine.InsertOne(t)
}

func GetAllType() (*[]Type, error) {
	t := make([]Type, 0)
	err := engine.Find(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

type Types struct {
	Types map[string]int64
	sync.Mutex
}

var types Types

// 添加type到types
func (t *Type) AddTypes() error {
	types.Lock()
	defer types.Unlock()
	if err := writeTypes(); err != nil {
		return err
	}
	types.Types[t.Name] = t.Id
	return nil
}

// 写入Types
func writeTypes() error {
	if types.Types == nil {
		ts, err := GetAllType()
		if err != nil {
			return err
		}
		types.Types = make(map[string]int64)
		for _, v := range *ts {
			types.Types[v.Name] = v.Id
		}
	}
	return nil
}

// 根据name获取id
func GetIdByName(name string) (int64, error) {
	types.Lock()
	defer types.Unlock()
	if err := writeTypes(); err != nil {
		return 0, err
	}
	if id, ok := types.Types[name]; ok {
		return id, nil
	} else {
		return 0, nil
	}
}
