package main

import (
	"net"
	"fmt"
	"io"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	l, err := net.Listen("tcp", ":5656")
	defer l.Close()
	CheckErr(err)
	for {
		conn, err := l.Accept()
		CheckErr(err)

		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			return //出错后返回
		}
		fmt.Println(string(buf[0:n]))
	}
}
