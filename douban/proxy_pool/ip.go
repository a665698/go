package proxy_pool

import (
	"douban/common"
	"net/http"
	"time"
)

var pendingIp chan string

func init() {
	pendingIp = make(chan string)
}

// 写入待处理IP
func WritePendingIP(ip string) {
	pendingIp <- ip
}

// 读取待处理IP
func ReadPendingIp() string {
	return <-pendingIp
}

// 删除代理IP
func DelIp(ip string) {
	common.DelIp(ip)
}

// 获取代理IP
func GetIp() string {
	for {
		ip, err := common.GetIp()
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		return ip
	}
}

// 判断待处理的IP是否可用
func CheckIp() {
	for {
		ip := ReadPendingIp()
		res, err := common.GetHttpRes("https://www.baidu.com", ip)
		if err != nil {
			continue
		}
		if res.StatusCode != http.StatusOK {
			continue
		}
		res.Body.Close()
		common.AddIpPool(ip)
	}
}
