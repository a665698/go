package main

import "encoding/json"

var HttpResultArr = map[int]string{
	0: "成功",
	1: "用户名密码错误",
	2: "链接错误",
}

type HttpResult struct {
	id int
	info string
}

func GetHttpResult(k int) []byte {
	info := HttpResult{id: k}
	if v, ok := HttpResultArr[k]; ok {
		info.info = v
	} else {
		info.info = "未知错误"
	}
	str, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	return str
}


