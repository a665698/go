package model

type ReleaseDate struct {
	Id         int64
	MovieId    int64
	Time       int64
	Remark     string
	CreateTime int64 `xorm:"created"`
}

func NewReleaseDate() *ReleaseDate {
	return &ReleaseDate{}
}

func (r *ReleaseDate) Insert() (int64, error) {
	return engine.InsertOne(r)
}
