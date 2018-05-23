package model

type MovieLanguage struct {
	Id         int64
	MovieId    int64
	LanguageId int64
	CreateTime int64 `xorm:"created"`
}

func NewMovieLanguage(movieId, languageId int64) *MovieLanguage {
	return &MovieLanguage{
		MovieId:    movieId,
		LanguageId: languageId,
	}
}

func (ml *MovieLanguage) Insert() (int64, error) {
	return engine.InsertOne(ml)
}
