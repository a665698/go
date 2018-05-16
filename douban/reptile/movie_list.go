package reptile

import (
	"douban/model"
	"encoding/json"
)

type MovieList struct {
	Movies []*model.Movie `json:"subjects"`
}

func NewMovieList() *MovieList {
	return &MovieList{}
}

// 获取movie列表
func getMovieList(tag string) {
	body, err := getHttp(MOVIE_LIST + tag)
	if err != nil {
		noticeLog(err)
	}
	movie := NewMovieList()
	err = json.Unmarshal(body, &movie)
	if err != nil {
		noticeLog(err)
	}
	noticeLog(movie.Movies[0])
	noticeLog(movie.Movies[1])
	//GetMainMovieQueue().Put(movie.Movies...)
}
