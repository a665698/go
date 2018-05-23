package model

import "sync"

type District struct {
	Id         int64
	Name       string
	CreateTime int64 `xorm:"created"`
}

func NewDistrict(name string) *District {
	return &District{
		Name: name,
	}
}

func (d *District) Insert() (int64, error) {
	return engine.InsertOne(d)
}

// 获取所有district
func GetAllDistrict() (*[]District, error) {
	districts := make([]District, 0)
	err := engine.Find(&districts)
	if err != nil {
		return nil, err
	}
	return &districts, nil
}

type Districts struct {
	Districts map[string]int64
	sync.Mutex
}

var districts Districts

func (d *District) AddDistricts() error {
	districts.Lock()
	defer districts.Unlock()
	if err := writeDistricts(); err != nil {
		return err
	}
	districts.Districts[d.Name] = d.Id
	return nil
}

// 写入districts
func writeDistricts() error {
	if districts.Districts == nil {
		ds, err := GetAllDistrict()
		if err != nil {
			return err
		}
		districts.Districts = make(map[string]int64)
		for _, v := range *ds {
			districts.Districts[v.Name] = v.Id
		}
	}
	return nil
}

// 根据name获取districtId
func GetDistrictIdByName(name string) (int64, error) {
	districts.Lock()
	defer districts.Unlock()
	if err := writeDistricts(); err != nil {
		return 0, err
	}
	if id, ok := districts.Districts[name]; ok {
		return id, nil
	} else {
		return 0, nil
	}
}
