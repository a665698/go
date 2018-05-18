package model

import "sync"

type Performer struct {
	Id         int64
	Name       string
	CreateTime int64 `xorm:"created"`
}

var performers Performers

type Performers struct {
	performer *[]Performer
	sync.Mutex
}

func NewPerformer() *Performer {
	return &Performer{}
}

func (p *Performer) Insert() (int64, error) {
	return engine.Insert(p)
}

func GetAllPerformer() (*[]Performer, error) {
	performers := make([]Performer, 0)
	err := engine.Find(&performers)
	if err != nil {
		return nil, err
	}
	return &performers, nil
}

func GetPerformerId(name string) (int64, error) {
	performers.Lock()
	defer performers.Unlock()
	if performers.performer == nil {
		var err error
		performers.performer, err = GetAllPerformer()
		if err != nil {
			return 0, err
		}
	}
	var performer Performer
	for _, performer = range *performers.performer {
		if performer.Name == name {
			return performer.Id, nil
		}
	}
	performer.Id = 0
	performer.Name = name
	id, err := performer.Insert()
	if err != nil {
		return 0, err
	}
	performer.Id = id
	*performers.performer = append(*performers.performer, performer)
	return id, nil
}
