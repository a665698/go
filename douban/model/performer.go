package model

import (
	"sync"
)

type Performer struct {
	Id         int64
	Name       string
	Type       uint8
	CreateTime int64 `xorm:"created"`
}

// 初始化
func NewPerformer(name string, t uint8) *Performer {
	return &Performer{
		Name: name,
		Type: t,
	}
}

// 添加演员记录
func (p *Performer) Insert() (int64, error) {
	return engine.InsertOne(p)
}

// 修改演员记录
func (p *Performer) Update() (int64, error) {
	return engine.Where("id = ? ", p.Id).Cols("type").Update(p)
}

// 获取所有演员列表
func GetAllPerformer() (*[]Performer, error) {
	performers := make([]Performer, 0)
	err := engine.Find(&performers)
	if err != nil {
		return nil, err
	}
	return &performers, nil
}

// 所有演员
type Performers struct {
	performer map[string]*Performer
	sync.Mutex
}

var performers Performers

// 添加演员到所有演员列表
func (p *Performer) AddPerformers() error {
	performers.Lock()
	defer performers.Unlock()
	err := writePerformers()
	if err != nil {
		return err
	}
	performers.performer[p.Name] = p
	return nil
}

// 初始化所有演员列表
func writePerformers() error {
	if performers.performer == nil {
		ps, err := GetAllPerformer()
		if err != nil {
			return err
		}
		performers.performer = make(map[string]*Performer)
		for k := range *ps {
			performers.performer[(*ps)[k].Name] = &(*ps)[k]
		}
	}
	return nil
}

// 获取演员详情
func GetPerformerByName(name string) (*Performer, error) {
	performers.Lock()
	defer performers.Unlock()
	err := writePerformers()
	if err != nil {
		return nil, err
	}
	if performer, ok := performers.performer[name]; ok {
		return performer, nil
	} else {
		return nil, nil
	}
}
