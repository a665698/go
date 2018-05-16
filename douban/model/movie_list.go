package model

type Movie struct {
	Id         int64  `json:"-"`
	MovieId    int64  `json:"id,string"`
	Rate       string `json:"rate"`
	Title      string `json:"title"`
	Cover      string `json:"cover"`
	CreateTime int64  `xorm:"created",json:"-"`
}

func NewMovieList() *Movie {
	return &Movie{}
}

func (m *Movie) GetAll() (*[]Movie, error) {
	movies := make([]Movie, 0)
	err := engine.Find(&movies)
	if err != nil {
		return nil, err
	}
	return &movies, nil
}
