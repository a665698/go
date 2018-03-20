package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5656")
	checkErr(err)
	defer conn.Close()

	input := bufio.NewReader(os.Stdin)
	fmt.Println("请输入1：")
	s, err := input.ReadString('\n')
	checkErr(err)
	conn.Write([]byte(strings.TrimSpace(s)))
	fmt.Println("请输入2：")
	s, err = input.ReadString('\n')
	checkErr(err)
	conn.Write([]byte(s))
}


