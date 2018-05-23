package model

import "sync"

type Language struct {
	Id         int64
	Name       string
	CreateTime int64 `xorm:"created"`
}

func NewLanguage(name string) *Language {
	return &Language{
		Name: name,
	}
}

// 写入一条到数据库
func (l *Language) Insert() (int64, error) {
	return engine.InsertOne(l)
}

// 获取所有language
func GetAllLanguage() (*[]Language, error) {
	ls := make([]Language, 0)
	if err := engine.Find(&ls); err != nil {
		return nil, err
	} else {
		return &ls, nil
	}
}

type Languages struct {
	Language map[string]int64
	sync.Mutex
}

var languages Languages

// 初始化language
func writeLanguages() error {
	if languages.Language == nil {
		ls, err := GetAllLanguage()
		if err != nil {
			return err
		}
		languages.Language = make(map[string]int64)
		for _, v := range *ls {
			languages.Language[v.Name] = v.Id
		}
	}
	return nil
}

// 添加到languages
func (l *Language) AddLanguages() error {
	languages.Lock()
	defer languages.Unlock()
	if err := writeLanguages(); err != nil {
		return err
	}
	languages.Language[l.Name] = l.Id
	return nil
}

// 根据name获取languageId
func GetLanguageIdByName(name string) (int64, error) {
	languages.Lock()
	defer languages.Unlock()
	if err := writeLanguages(); err != nil {
		return 0, err
	}
	if id, ok := languages.Language[name]; ok {
		return id, nil
	} else {
		return 0, nil
	}
}
