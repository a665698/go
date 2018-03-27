package common

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"time"
	"errors"
	"bytes"
	"mime/multipart"
	"os"
	"io"
	"strconv"
	"fmt"
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

	Code
}

type Code struct {
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
		return nil,errors.New("错误码：" + strconv.Itoa(s.ErrCode) + ",错误信息：" + s.ErrMsg)
	}
	return s, nil
}

// 发送post请求
func Post(url string, postInfo []byte) error {
	result, err := http.Post(url, "application/json; encoding=utf-8", strings.NewReader(string(postInfo)))
	if err != nil {
		return err
	}
	return postReturnValid(result.Body)
}

// 上传文件
func UploadFile(url, fileName string) error {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("media", fileName)
	if err != nil {
		return err
	}
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := io.Copy(fw, f); err != nil {
		return err
	}
	fmt.Println(url)
	req, err := http.NewRequest("post", url, &b)
	if err != nil {
		return err
	}
	fmt.Println(b)
	fmt.Println(w.FormDataContentType())
	req.Header.Set("Content-Type", w.FormDataContentType())
	h := &http.Client{}
	res, err := h.Do(req)
	if err != nil {
		return err
	}
	return postReturnValid(res.Body)
}

// post请求结果检验
func postReturnValid(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	r := Code{}
	if err = json.Unmarshal(b, &r); err != nil {
		return err
	}
	if r.ErrCode != 0 {
		return errors.New("错误码：" + strconv.Itoa(r.ErrCode) + ",错误信息：" + r.ErrMsg)
	}
	return nil
}


