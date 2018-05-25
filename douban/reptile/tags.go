package reptile

import (
	"douban/common"
	"encoding/json"
	"time"
)

type Tags struct {
	Tag []string `json:"tags"`
}

func tickTags() {
	go movieInfo()
	for {
		common.DelMovieId()
		getTags()
		time.Sleep(time.Hour * 2)
	}
}

// 获取标签列表
func getTags() {
	common.NoticeLog("start get tags")
	body, err := common.GetHttp(common.TAGS_URL, "")
	if err != nil {
		common.NoticeLog(err)
		return
	}
	tags := Tags{}
	err = json.Unmarshal(body, &tags)
	if err != nil {
		common.NoticeLog(err)
		return
	}
	tags.tagsHandle()
	common.NoticeLog("over tags")
}

// 处理标签列表
func (t *Tags) tagsHandle() {
	for _, v := range t.Tag {
		getMovieList(v)
	}
}
