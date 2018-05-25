package reptile

import (
	"douban/common"
	"douban/model"
	"douban/proxy_pool"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
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
	movieInfoChan := make(chan *MovieInfo)
	go MovieInfoHandle(movieInfoChan)
	for i := 0; i < 10; i++ {
		go GetMovieInfo(movieInfoChan)
	}
}

// 获取movie详情
func GetMovieInfo(movieInfo chan<- *MovieInfo) {
	for {
		movie := MoviePoll()
		// 是否存在待处理Id列表中,存在跳过
		b, _ := common.IsExistMovieId(movie.MovieId)
		if b == 1 {
			continue
		}
		// 是否存在已处理完成的列表中,存在跳过
		mInfo, err := model.GetMovieByMovieId(movie.MovieId)
		if err != nil {
			common.NoticeLog(err)
			continue
		}
		if mInfo != nil {
			continue
		}
		// 获取代理IP
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
		if !movie.MovieInfoHandle(dom, movieInfo) {
			common.NoticeLog(fmt.Sprintf("id: %d 页面解析错误", movie.MovieId))
			MoviePut(movie)
			continue
		}
	}
}

// 详情页面解析
func (m *MovieInfo) MovieInfoHandle(d *goquery.Document, movieInfo chan<- *MovieInfo) bool {
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
				m.Directed = v[1]
			case "编剧":
				m.Celebrity = v[1]
			case "类型":
				m.Type = v[1]
			case "制片国家/地区":
				m.District = v[1]
			case "语言":
				m.Language = v[1]
			case "上映日期":
				m.ReleaseDate = v[1]
			case "片长":
				m.Runtime = v[1]
			case "又名":
				m.Alias = v[1]
			}
		}
	}
	m.Summary = strings.TrimSpace(d.Find("#link-report").Text())
	star := d.Find(".ratings-on-weight .rating_per")
	if star.Length() == 5 {
		m.Star5 = star.Eq(0).Text()
		m.Star4 = star.Eq(1).Text()
		m.Star3 = star.Eq(2).Text()
		m.Star2 = star.Eq(3).Text()
		m.Star1 = star.Eq(4).Text()
	}
	//movieBytes, _ := json.Marshal(m)
	//common.AddMovieInfo(movieBytes)
	movieInfo <- m
	common.AddMovieId(m.MovieId)
	common.NoticeLog(m.Title + ":处理完成")
	return true
}

func MovieInfoHandle(movieInfo <-chan *MovieInfo) {
	for {
		//info, err := common.GetMovieInfo()
		//if err != nil {
		//	common.NoticeLog(err)
		//	return
		//	//time.Sleep(time.Second * 10)
		//}
		//movie := NewMovieInfo()
		//err = json.Unmarshal(info, movie)
		//if err != nil {
		//	common.NoticeLog(err)
		//	continue
		//}
		movie := <-movieInfo
		mInfo, err := model.GetMovieByMovieId(movie.MovieId)
		if err != nil {
			common.NoticeLog(err)
			continue
		}
		if mInfo != nil {
			continue
		}
		m := model.NewMovie()
		m.Title, m.MovieId, m.Rate, m.Cover = movie.Title, movie.MovieId, movie.Rate, movie.Cover
		m.Star1, m.Star2, m.Star3, m.Star4, m.Star5 = movie.Star1, movie.Star2, movie.Star3, movie.Star4, movie.Star5
		_, err = m.Insert()
		if err != nil {
			common.NoticeLog(err)
			continue
		}
		movie.DirectedHandle(m.Id, 1)
		movie.DirectedHandle(m.Id, 0)
		movie.TypeHandle(m.Id)
		movie.DistrictHandle(m.Id)
		movie.LanguageHandle(m.Id)
		movie.ReleaseDateHandle(m.Id)
		movie.RuntimeHandle(m.Id)
		movie.AliasHandle(m.Id)
		movie.SummaryHandle(m.Id)
	}
}

// movie演员处理
// id: movie主键ID
// b: {1: Directed, 0: Celebrity}
func (m *MovieInfo) DirectedHandle(id int64, b uint8) {
	var directed []string
	if b == 1 {
		directed = strings.Split(m.Directed, "/")
	} else {
		directed = strings.Split(m.Celebrity, "/")
	}
	for _, v := range directed {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		performer, err := model.GetPerformerByName(v)
		if err != nil {
			common.NoticeLog(err)
			continue
		}
		if performer == nil {
			performer = model.NewPerformer(v, b)
			_, err = performer.Insert()
			if err != nil {
				common.NoticeLog(err)
				continue
			}
			performer.AddPerformers()
		}
		if b == 1 && performer.Type != 1 {
			performer.Type = 1
			performer.Update()
		}
		if b == 0 {
			i := strings.Index(m.Directed, v)
			if i >= 0 {
				continue
			}
		}
		_, err = model.NewMoviePerformer(id, performer.Id).Insert()
		if err != nil {
			common.NoticeLog(err)
		}
	}
}

// movie type处理
func (m *MovieInfo) TypeHandle(id int64) {
	t := strings.Split(m.Type, "/")
	for _, v := range t {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		typeId, err := model.GetIdByName(v)
		if err != nil {
			common.NoticeLog(err)
			continue
		}
		if typeId == 0 {
			Type := model.NewType(v)
			_, err := Type.Insert()
			if err != nil {
				common.NoticeLog(err)
				continue
			}
			Type.AddTypes()
			typeId = Type.Id
		}
		_, err = model.NewMovieType(id, typeId).Insert()
		if err != nil {
			common.NoticeLog(err)
		}
	}
}

// movie district处理
func (m *MovieInfo) DistrictHandle(id int64) {
	s := strings.Split(m.District, "/")
	for _, v := range s {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		districtId, err := model.GetDistrictIdByName(v)
		if err != nil {
			common.NoticeLog(err)
			continue
		}
		if districtId == 0 {
			district := model.NewDistrict(v)
			_, err := district.Insert()
			if err != nil {
				common.NoticeLog(err)
				continue
			}
			district.AddDistricts()
			districtId = district.Id
		}
		_, err = model.NewMovieDistrict(id, districtId).Insert()
		if err != nil {
			common.NoticeLog(err)
		}
	}
}

// movie language处理
func (m *MovieInfo) LanguageHandle(id int64) {
	s := strings.Split(m.Language, "/")
	for _, v := range s {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		languageId, err := model.GetLanguageIdByName(v)
		if err != nil {
			common.NoticeLog(err)
			continue
		}
		if languageId == 0 {
			language := model.NewLanguage(v)
			_, err := language.Insert()
			if err != nil {
				common.NoticeLog(err)
				continue
			}
			language.AddLanguages()
			languageId = language.Id
		}
		_, err = model.NewMovieLanguage(id, languageId).Insert()
		if err != nil {
			common.NoticeLog(err)
		}
	}
}

// movie release_date处理
func (m *MovieInfo) ReleaseDateHandle(id int64) {
	s := strings.Split(m.ReleaseDate, "/")
	tl, _ := time.LoadLocation("Asia/Shanghai")
	for _, v := range s {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		fIndex := strings.IndexRune(v, '(')
		releaseDate := model.NewReleaseDate()
		releaseDate.MovieId = id
		if fIndex > 0 {
			replacer := strings.NewReplacer("(", "", ")", "")
			releaseDate.Remark = replacer.Replace(v[fIndex:])
			v = v[:fIndex]
		}
		var t time.Time
		if _, err := strconv.Atoi(v); err != nil {
			t, _ = time.ParseInLocation("2006-01-02", v, tl)
		} else {
			t, _ = time.ParseInLocation("2006", v, tl)
		}
		releaseDate.Time = t.Unix()
		if _, err := releaseDate.Insert(); err != nil {
			common.NoticeLog(err)
		}
	}
}

// movie runtime处理
func (m *MovieInfo) RuntimeHandle(id int64) {
	s := strings.Split(m.Runtime, " /")
	for k := range s {
		str := strings.Split(s[k], "/")
		in := strings.IndexRune(s[k], '(')
		s[k] = str[0] + ")"
		for i := 1; i < len(str); i++ {
			s = append(s, str[k][:in]+"("+str[i])
		}
	}
	for _, v := range s {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		fIndex := strings.IndexRune(v, '(')
		var districtId int64
		var err error
		if fIndex > 0 {
			districtName := v[fIndex+1 : len(v)-1]
			districtId, err = model.GetDistrictIdByName(districtName)
			if err != nil {
				common.NoticeLog(err)
				continue
			}
			if districtId == 0 {
				district := model.NewDistrict(districtName)
				_, err := district.Insert()
				if err != nil {
					common.NoticeLog(err)
					continue
				}
				district.AddDistricts()
				districtId = district.Id
			}
		}
		reg, _ := regexp.Compile("^[0-9]+")
		str := reg.FindStringSubmatch(v)
		var t int
		if len(str) > 0 {
			t, _ = strconv.Atoi(str[0])
		}
		_, err = model.NewRuntime(t, districtId, id).Insert()
		if err != nil {
			common.NoticeLog(err)
		}
	}
}

// movie alias 处理
func (m *MovieInfo) AliasHandle(id int64) {
	s := strings.Split(m.Alias, "/")
	for _, v := range s {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		_, err := model.NewAlias(id, v).Insert()
		if err != nil {
			common.NoticeLog(err)
		}
	}
}

// movie summary处理
func (m *MovieInfo) SummaryHandle(id int64) {
	str := "(展开全部)"
	fIndex := strings.Index(m.Summary, str)
	if fIndex > 0 {
		m.Summary = m.Summary[fIndex+len(str):]
	}
	m.Summary = strings.TrimSpace(m.Summary)
	m.Summary = strings.TrimSuffix(m.Summary, "©豆瓣")
	arr := []rune(m.Summary)
	var currentArr []rune
	var isDel, upIsSpace bool
	for _, v := range arr {
		if unicode.IsSpace(v) {
			if !isDel && !upIsSpace {
				currentArr = append(currentArr, v)
			}
			upIsSpace = true
			continue
		}
		upIsSpace = false
		currentArr = append(currentArr, v)
		isDel = true
		if (v >= 65 && v <= 106) || (v >= 113 && v <= 122) {
			isDel = false
		}
	}
	m.Summary = string(currentArr)
	_, err := model.NewSummary(id, m.Summary).Insert()
	if err != nil {
		common.NoticeLog(err)
	}
}
