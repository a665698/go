package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main() {
	file, err := os.OpenFile("te.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	f := ""
	for {
		line, err := buf.ReadString('\n')
		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		fmt.Println(line)
		f += line
	}
	os.Truncate("te.txt", 0)
	file.Write([]byte(f))
}



