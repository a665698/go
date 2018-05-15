package reptile

import (
	"encoding/json"
	"time"
)

type MovieList struct {
	Movies []Movie `json:"subjects"`
}

type Movie struct {
	Rate     string `json:"rate"`
	CoverX   int    `json:"cover_x"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Playable bool   `json:"playable"`
	Cover    string `json:"cover"`
	Id       string `json:"id"`
	CoverY   int    `json:"cover_y"`
	IsNew    bool   `json:"is_new"`
}

func NewMovie() *Movie {
	return &Movie{}
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
	noticeLog(movie)
}

func getCurrentMovieList() ([]Movie, error) {
	rows, err := mySql.Query("select movie_id,rate,title,cover,is_new, from lnn_movie_list")
	if err != nil {
		return nil, err
	}
	list := make([]Movie, 0)
	var (
		isNew uint8
		id    int
		rate  string
		title string
		cover string
	)
	var movie Movie
	for rows.Next() {
		if err = rows.Scan(&id, &rate, &title, &cover, &isNew); err != nil {
			return nil, err
		}
		movie.Id, movie.Rate, movie.Title, movie.Cover = string(id), rate, title, cover
		if isNew == 1 {
			movie.IsNew = true
		} else {
			movie.IsNew = false
		}
		list = append(list, movie)
	}
	return list, nil
}

func movieHandle() {
	tx, err := mySql.Begin()
	if err != nil {
		noticeLog(err)
		return
	}
	stmt, err := tx.Prepare("insert into lnn_movie_list (`movie_id`,`rate`,`title`,`cover`,`is_new`,`create_time`) values(?,?,?,?,?,?)")
	if err != nil {
		noticeLog(err)
		return
	}
	var isNew uint8
	createTime := time.Now().Unix()
}
