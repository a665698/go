package reptile

import (
	"douban/common"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"testing"
)

func TestGetMovieInfo(t *testing.T) {
	res, err := common.GetHttpRes("https://movie.douban.com/subject/1866354/", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	dom, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		fmt.Println(err)
		return
	}
	movie := NewMovieInfo()
	movie.Title = "测试1"
	movie.MovieId = 111
	movie.Rate = "8.2"
	movie.Cover = "http://www.aaa.com"
	movieInfoHandle(dom, movie)
}

func TestMovieInfoHandle(t *testing.T) {
	MovieInfoHandle()
}

func TestMovieInfo_RuntimeHandle(t *testing.T) {
	info, _ := common.GetMovieInfo()
	movie := NewMovieInfo()
	json.Unmarshal(info, movie)
	movie.RuntimeHandle(1545646)
}

func TestMovieInfo_SummaryHandle(t *testing.T) {
	info, _ := common.GetMovieInfo()
	movie := NewMovieInfo()
	json.Unmarshal(info, movie)
	movie.SummaryHandle(1545646)
}
