package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"
	"transmit/log"
)

type Server struct {
	Addr string
	Conn net.Conn
	Log  io.Writer
}

func NewServer(addr string, conn net.Conn) *Server {
	l, err := os.OpenFile(addr+".log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Error(err)
	} else {
		l = os.Stdout
	}
	return &Server{
		Addr: addr,
		Conn: conn,
		Log:  l,
	}
}

func Main() {
	listen, err := net.Listen("tcp", ":5566")
	if err != nil {
		log.Panic(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Error(err)
			continue
		}
		log.Notify("收到tcp连接")
		go ConnHandle(conn)
	}
}

func ConnHandle(conn net.Conn) {
	defer conn.Close()
	buf := bufio.NewReader(conn)
	for {
		str, err := buf.ReadString('\n')
		if err != nil {
			break
		}
		conn.Write(CommandHandle(str))
	}
}

func CommandHandle(command string) []byte {
	c := strings.Split(strings.TrimSpace(command), " ")
	var cmd *exec.Cmd
	if len(c) > 1 {
		cmd = exec.Command(c[0], c[1:]...)
	} else {
		cmd = exec.Command(c[0])
	}
	out, err := cmd.StdoutPipe()
	if err != nil {
		return []byte(err.Error())
	}
	err = cmd.Start()
	if err != nil {
		return []byte(err.Error())
	}
	reader := bufio.NewReader(out)
	var b []byte
	var buf = make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				return []byte(err.Error())
			}
			break
		}
		b = append(b, buf[:n]...)
	}
	cmd.Wait()
	fmt.Println(96, string(b))
	return b
}
