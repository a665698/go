package socket

import "encoding/json"

var HttpResultArr = map[int]string{
	0: "成功",
	1: "用户名密码错误",
	2: "链接错误",
}

type HttpResult struct {
	Status int `json:"status"`
	Info string `json:"info"`
}

func GetHttpResult(k int) []byte {
	info := HttpResult{Status: k}
	if v, ok := HttpResultArr[k]; ok {
		info.Info = v
	} else {
		info.Info = "未知错误"
	}
	str, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	return str
}


