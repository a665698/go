package socket

import (
	"github.com/gorilla/websocket"
	"errors"
	"fmt"
)

type client struct {
	conn *websocket.Conn
	name string
	flag string
}

var clients = make(map[string]*client)
var addUser = make(chan *client)
var delUser = make(chan *client)

func init() {
	go channelUser()
}

// 获取所有用户
func GetClients() map[string]*client {
	return clients
}

// 获取用户
func GetUser(key string) (*client, error) {
	if val, ok := clients[key]; !ok {
		return nil, errors.New("用户不存在")
	} else {
		return val, nil
	}
}

// 新建用户
func NewUser(flag, name string) {
	user := &client{name: name, flag: flag}
	addUser <- user
}

// 用户退出
func (c *client) QuitUser() {
	delUser <- c
}

// 用户通道,添加和删除用户
func channelUser() {
	for {
		select {
		case user := <- addUser:
			clients[user.flag] = user
		case user := <- delUser:
			delete(clients, user.flag)
		}
	}
}

// 读取用户发送的消息
func (c *client) readMes() {
	for {
		mt, message, err := c.conn.ReadMessage()
		if err != nil {
			if mt == -1 {
				sendSysMessage(c.name + " 退出")
				c.QuitUser()
			}
			fmt.Println("消息获取失败: ", err, "\n消息类型", mt)
			break
		}
		fmt.Println("收到消息: ", string(message), " \n消息类型: ", mt)
		sendTextMessage(string(message), c.name)
	}
}
