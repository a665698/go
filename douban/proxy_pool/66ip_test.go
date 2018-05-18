package proxy_pool

import (
	"testing"
)

func TestGet66Ip(t *testing.T) {
	go CheckIp()
	Get66Ip()
}
