package model

import (
	"sync"
)

type Movie struct {
	Id         int64
	MovieId    int64
	Title      string
	Rate       string
	Cover      string
	Star5      string
	Star4      string
	Star3      string
	Star2      string
	Star1      string
	CreateTime int64 `xorm:"created"`
}

func NewMovie() *Movie {
	return &Movie{}
}

func (m *Movie) Insert() (int64, error) {
	return engine.InsertOne(m)
}

func (m *Movie) Del() (int64, error) {
	return engine.Delete(m)
}

func GetAllMovie() (*[]Movie, error) {
	movies := make([]Movie, 0)
	err := engine.Find(&movies)
	if err != nil {
		return nil, err
	}
	return &movies, nil
}

type Movies struct {
	Movies map[int64]*Movie
	sync.Mutex
}

var movies Movies

// 获取movie
func GetMovieByMovieId(movieId int64) (*Movie, error) {
	movies.Lock()
	defer movies.Unlock()
	if movies.Movies == nil {
		m, err := GetAllMovie()
		if err != nil {
			return nil, err
		}
		movies.Movies = make(map[int64]*Movie)
		for _, v := range *m {
			movies.Movies[v.MovieId] = &v
		}
	}
	if m, ok := movies.Movies[movieId]; ok {
		return m, nil
	} else {
		return nil, nil
	}
}

// 获取重复数据
func GetRepeat() (*[]Movie, error) {
	movies := make([]Movie, 0)
	err := engine.Select("COUNT(id) as cid,id").GroupBy("movie_id").Having("cid >1").Find(&movies)
	if err != nil {
		return nil, err
	}
	return &movies, nil
}

func GetList() (*[]Movie, error) {
	ms := make([]Movie, 0)
	err := engine.Desc("rate").Find(&ms)
	if err != nil {
		return nil, err
	}
	return &ms, nil
}
