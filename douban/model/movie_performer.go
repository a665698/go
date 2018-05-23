package model

type MoviePerformer struct {
	Id          int64
	MovieId     int64
	PerformerId int64
	CreateTime  int64 `xorm:"created"`
}

func NewMoviePerformer(movieId, performerId int64) *MoviePerformer {
	return &MoviePerformer{
		MovieId:     movieId,
		PerformerId: performerId,
	}
}

func (mp *MoviePerformer) Insert() (int64, error) {
	return engine.InsertOne(mp)
}
