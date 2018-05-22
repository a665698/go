package proxy_pool

import (
	"fmt"
	"testing"
)

func TestGetIp(t *testing.T) {
	ip := GetIp()
	fmt.Println(ip)
}
