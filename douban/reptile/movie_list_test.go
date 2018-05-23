package reptile

import (
	"douban/common"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"testing"
)

func TestGetMovieInfo(t *testing.T) {
	res, err := common.GetHttpRes("https://movie.douban.com/subject/27018285/", "")
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
