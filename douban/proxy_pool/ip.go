package proxy_pool

import (
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type IpPool struct {
	pendingIp chan string
	ip        []string
	sync.RWMutex
}

var ipPool *IpPool

func init() {
	ipPool = &IpPool{
		pendingIp: make(chan string),
		ip:        make([]string, 0),
	}
}

// 写入待处理IP
func WritePendingIP(ip string) {
	ipPool.pendingIp <- ip
}

// 读取待处理IP
func ReadPendingIp() string {
	return <-ipPool.pendingIp
}

// 添加ip
func (i *IpPool) Add(ip string) {
	i.Lock()
	defer i.Unlock()
	i.ip = append(i.ip, ip)
}

// 获取一条IP,如果不存在则每秒执行一次，直到获取到为止
func GetIp() (int, string) {
	for {
		ipPool.RLock()
		l := len(ipPool.ip)
		if l <= 0 {
			ipPool.RUnlock()
			time.Sleep(time.Second)
			continue
		}
		rand.Seed(time.Now().Unix())
		k := rand.Intn(l) - 1
		val := ipPool.ip[k]
		ipPool.RUnlock()
		return k, val
	}
}

// 判断待处理的IP是否可用
func CheckIp() {
	for {
		ip := ReadPendingIp()
		res, err := GetHttp("https://www.baidu.com", ip)
		if err != nil {
			continue
		}
		if res.StatusCode != http.StatusOK {
			continue
		}
		ipPool.Add(ip)
	}
}
