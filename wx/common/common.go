package common

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strings"
	"time"
	"errors"
)

type GetResponse struct {
	// token
	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`

	// 用户
	Subscribe int8 `json:"subscribe"`
	Openid string `json:"openid"`
	Nickname string `json:"nickname"`
	Sex int8 `json:"sex"`
	Language string `json:"language"`
	City string `json:"city"`
	Province string `json:"province"`
	Country string `json:"country"`
	HeadImgUrl string `json:"headimgurl"`
	SubscribeTime time.Duration `json:"subscribe_time"`
	UnionId string `json:"unionid"`
	Remark string `json:"remark"`
	GroupId int8 `json:"groupid"`
	TagIdList []int `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene int `json:"qr_scene"`
	QrSceneStr string `json:"qr_scene_str"`

	// 错误
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

// 发送GET请求获取数据
func Get(url string) (*GetResponse, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	r, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil ,err
	}
	s := &GetResponse{}
	err = json.Unmarshal(r, s)
	if err != nil {
		return nil, err
	}
	if s.ErrCode != 0 {
		return nil,errors.New("错误码：" + string(s.ErrCode) + ",错误信息：" + s.ErrMsg)
	}
	return s, nil
}

func Post(url, postInfo string) bool {
	result, err := http.Post(url, "application/json; encoding=utf-8", strings.NewReader(postInfo))
	if err != nil {
		fmt.Println(err)
		return false
	}
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	var r map[string]interface{}
	if err = json.Unmarshal(body, &r); err != nil {
		fmt.Println(err)
		return false
	}
	if errCode, ok := r["errcode"]; !ok {
		fmt.Println("errcode not fount")
		return false
	} else {
		code := int(errCode.(float64))
		if code == 0 {
			return true
		} else {
			//fmt.Println(GetError(code))
			return false
		}
	}
}
