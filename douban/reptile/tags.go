package reptile

import (
	"encoding/json"
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
	getMovieList(t.Tag[0])
	//for _, v := range t.Tag {
	//	go getMovieList(v)
	//}
}
