package model

type MovieType struct {
	Id         int64
	MovieId    int64
	TypeId     int64
	CreateTime int64 `xorm:"created"`
}

func NewMovieType(movieId, typeId int64) *MovieType {
	return &MovieType{
		MovieId: movieId,
		TypeId:  typeId,
	}
}

func (my *MovieType) Insert() (int64, error) {
	return engine.InsertOne(my)
}
