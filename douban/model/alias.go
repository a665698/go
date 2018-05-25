package model

type Alias struct {
	Id         int64
	MovieId    int64
	Name       string
	CreateTime int64 `xorm:"created"`
}

func NewAlias(movieId int64, name string) *Alias {
	return &Alias{
		MovieId: movieId,
		Name:    name,
	}
}

func (a *Alias) Insert() (int64, error) {
	return engine.InsertOne(a)
}

func (a *Alias) Del() (int64, error) {
	return engine.Delete(a)
}

func DelAliasByMovieId(id int64) (int64, error) {
	a := new(Alias)
	a.MovieId = id
	return a.Del()
}
