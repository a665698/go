package reptile

import (
	"douban/model"
	"encoding/json"
	"strconv"
	"time"
)

type MovieList struct {
	Movies []*model.Movie `json:"subjects"`
}

func NewMovieList() *MovieList {
	return &MovieList{}
}

// 获取movie列表
func getMovieList(tag string) {
	NoticeLog("start tag:" + tag)
	body, err := getHttp(MOVIE_LIST + tag)
	if err != nil {
		NoticeLog(err)
		return
	}
	movie := NewMovieList()
	err = json.Unmarshal(body, &movie)
	if err != nil {
		NoticeLog(err)
		return
	}
	MoviePut(movie.Movies...)
	NoticeLog("over tag:" + tag)
}

//
func movieListHandle() {
	currentMovie, err := model.GetAllMovie()
	if err != nil {
		NoticeLog(err)
		return
	}
	var isAdd bool
	addMovie := make([]*model.Movie, 0)
	for {
		isAdd = true
		movie := MoviePoll()
		if movie == nil {
			if len(addMovie) > 0 {
				row, err := model.AddMovie(addMovie)
				if err != nil {
					NoticeLog(err)
				} else {
					NoticeLog("add row:" + strconv.FormatInt(row, 10))
				}
				addMovie = nil
				addMovie = make([]*model.Movie, 0)
			}
			time.Sleep(time.Second * 10)
			continue
		}
		for _, val := range *currentMovie {
			if movie.MovieId == val.MovieId {
				isAdd = false
				break
			}
		}
		if isAdd {
			NoticeLog("add movie title:" + movie.Title)
			addMovie = append(addMovie, movie)
			*currentMovie = append(*currentMovie, *movie)
		}
	}
}
