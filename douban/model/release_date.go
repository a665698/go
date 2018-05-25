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

func (r *ReleaseDate) Del() (int64, error) {
	return engine.Delete(r)
}

func DelReleaseDateByMovieId(id int64) (int64, error) {
	rd := new(ReleaseDate)
	rd.MovieId = id
	return rd.Del()
}
