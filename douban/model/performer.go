package model

import "sync"

type Performer struct {
	Id         int64
	Name       string
	CreateTime int64 `xorm:"created"`
}

// 初始化
func NewPerformer() *Performer {
	return &Performer{}
}

// 添加演员记录
func (p *Performer) Insert() (int64, error) {
	id, err := engine.Insert(p)
	if err != nil {
		return 0, err
	}
	performers.performer[p.Name] = id
	return id, nil
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

type Performers struct {
	performer map[string]int64
	sync.Mutex
}

var performers *Performers

func init() {
	performers = &Performers{
		performer: make(map[string]int64),
	}
}

// 获取演员ID，不存在时添加
func GetPerformerId(name string) (int64, error) {
	performers.Lock()
	defer performers.Unlock()
	if len(performers.performer) == 0 {
		ps, err := GetAllPerformer()
		if err != nil {
			return 0, err
		}
		for _, p := range *ps {
			performers.performer[p.Name] = p.Id
			if p.Name == name {
			}
		}
	}
	if id, ok := performers.performer[name]; ok {
		return id, nil
	}
	var performer Performer
	performer.Name = name
	return performer.Insert()
}
