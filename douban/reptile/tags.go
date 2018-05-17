package reptile

import (
	"encoding/json"
	"time"
)

type Tags struct {
	Tag []string `json:"tags"`
}

func tickTags() {
	go movieListHandle()
	for {
		getTags()
		time.Sleep(time.Hour * 2)
	}
}

// 获取标签列表
func getTags() {
	noticeLog("start get tags")
	body, err := getHttp(TAGS_URL)
	if err != nil {
		noticeLog(err)
		return
	}
	tags := Tags{}
	err = json.Unmarshal(body, &tags)
	if err != nil {
		noticeLog(err)
		return
	}
	tags.tagsHandle()
	noticeLog("over tags")
}

// 处理标签列表
func (t *Tags) tagsHandle() {
	for _, v := range t.Tag {
		go getMovieList(v)
	}
}
