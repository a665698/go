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
	noticeLog("start tag:" + tag)
	body, err := getHttp(MOVIE_LIST + tag)
	if err != nil {
		noticeLog(err)
		return
	}
	movie := NewMovieList()
	err = json.Unmarshal(body, &movie)
	if err != nil {
		noticeLog(err)
		return
	}
	MoviePut(movie.Movies...)
	noticeLog("over tag:" + tag)
}

func movieListHandle() {
	currentMovie, err := model.GetAllMovie()
	if err != nil {
		noticeLog(err)
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
					noticeLog(err)
				} else {
					noticeLog("add row:" + strconv.FormatInt(row, 10))
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
			noticeLog("add movie title:" + movie.Title)
			addMovie = append(addMovie, movie)
			*currentMovie = append(*currentMovie, *movie)
		}
	}
}
