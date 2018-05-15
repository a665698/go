package reptile

import (
	"encoding/json"
	"sync"
)

type Tags struct {
	Tag []string `json:"tags"`
}

// 获取标签列表
func getTags() {
	body, err := getHttp(TAGS_URL)
	if err != nil {
		noticeLog(err)
	}
	tags := Tags{}
	err = json.Unmarshal(body, &tags)
	if err != nil {
		noticeLog(err)
	}
	tags.tagsHandle()
}

// 处理标签列表
func (t *Tags) tagsHandle() {
	for _, v := range t.Tag {
		go getMovieList(v)
	}
}
