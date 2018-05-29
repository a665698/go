package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"transmit/log"
)

func Client() {
	connect, err := net.Dial("tcp", "192.168.1.27:5566")
	if err != nil {
		log.Panic(err)
	}
	defer connect.Close()
	go func() {
		input := os.Stdin
		red := bufio.NewReader(input)
		for {
			str, err := red.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}
			connect.Write([]byte(str))
		}
	}()
	buf := bufio.NewReader(connect)
	for {
		str, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Print(str)
	}
}
