package reptile

import (
	"douban/common"
	"douban/proxy_pool"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

type MovieList struct {
	Movies []*MovieInfo `json:"subjects"`
}

func NewMovieList() *MovieList {
	return &MovieList{}
}

type MovieInfo struct {
	MovieId     int64  `json:"id,string"`
	Rate        string `json:"rate"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Directed    string `json:"directed"`
	Celebrity   string `json:"celebrity"`
	Type        string `json:"type"`
	District    string `json:"district"`
	Language    string `json:"language"`
	ReleaseDate string `json:"release_date"`
	Runtime     string `json:"runtime"`
	Alias       string `json:"alias"`
	Summary     string `json:"summary"`
	Star5       string `json:"star_5"`
	Star4       string `json:"star_4"`
	Star3       string `json:"star_3"`
	Star2       string `json:"star_2"`
	Star1       string `json:"star_1"`
}

var movieInfoChannel chan MovieInfo

func init() {
	movieInfoChannel = make(chan MovieInfo)
}

func NewMovieInfo() *MovieInfo {
	return &MovieInfo{}
}

// 获取movie列表,并添加到队列中
func getMovieList(tag string) {
	common.NoticeLog("start tag:" + tag)
	body, err := common.GetHttp(common.MOVIE_LIST+tag, "")
	if err != nil {
		common.NoticeLog(err)
		return
	}
	movie := NewMovieList()
	err = json.Unmarshal(body, &movie)
	if err != nil {
		common.NoticeLog(err)
		return
	}
	MoviePut(movie.Movies...)
	common.NoticeLog("over tag:" + tag)
}

func movieInfo() {
	for i := 0; i < 5; i++ {
		go GetMovieInfo()
	}
}

// 获取movie详情
func GetMovieInfo() {
	for {
		movie := MoviePoll()
		ip := proxy_pool.GetIp()
		res, err := common.GetHttpRes(common.MOVIE_INFO+strconv.FormatInt(movie.MovieId, 10), ip)
		if err != nil {
			common.NoticeLog(err)
			MoviePut(movie)
			proxy_pool.DelIp(ip)
			continue
		}
		if res.StatusCode != http.StatusOK {
			common.NoticeLog(fmt.Sprintf("MovieId: %d statusCode: %d proxyIp: %s", movie.MovieId, res.StatusCode, ip))
			MoviePut(movie)
			proxy_pool.DelIp(ip)
			continue
		}
		dom, err := goquery.NewDocumentFromResponse(res)
		if err != nil {
			common.NoticeLog(err)
			MoviePut(movie)
			continue
		}
		if !movieInfoHandle(dom, movie) {
			common.NoticeLog(fmt.Sprintf("id: %d 页面解析错误", movie.MovieId))
			MoviePut(movie)
			continue
		}
	}
}

// 详情页面解析
func movieInfoHandle(d *goquery.Document, movie *MovieInfo) bool {
	b, _ := common.IsExistMovieId(movie.MovieId)
	if b == 1 {
		return true
	}
	info := d.Find("#info")
	if info.Length() == 0 {
		return false
	}
	text := strings.Split(info.Text(), "\n")
	var v []string
	for _, val := range text {
		v = strings.Split(val, ":")
		if len(v) == 2 {
			switch strings.TrimSpace(v[0]) {
			case "导演":
				movie.Directed = v[1]
			case "编剧":
				movie.Celebrity = v[1]
			case "类型":
				movie.Type = v[1]
			case "制片国家/地区":
				movie.District = v[1]
			case "语言":
				movie.Language = v[1]
			case "上映日期":
				movie.ReleaseDate = v[1]
			case "片长":
				movie.Runtime = v[1]
			case "又名":
				movie.Alias = v[1]
			}
		}
	}
	movie.Summary = strings.TrimSpace(d.Find("#link-report").Text())
	star := d.Find(".ratings-on-weight .rating_per")
	if star.Length() == 5 {
		movie.Star5 = star.Eq(0).Text()
		movie.Star4 = star.Eq(1).Text()
		movie.Star3 = star.Eq(2).Text()
		movie.Star2 = star.Eq(3).Text()
		movie.Star1 = star.Eq(4).Text()
	}
	movieBytes, _ := json.Marshal(movie)
	common.AddMovieInfo(movieBytes)
	common.AddMovieId(movie.MovieId)
	common.NoticeLog(movie.Title + ":处理完成")
	return true
}
