package socket

import (
	"github.com/gorilla/websocket"
	"fmt"
	"sync"
)

type client struct {
	conn *websocket.Conn
	name string
	flag string
}

var clients = Clients{client: make(map[string]*client)}

type Clients struct {
	client map[string]*client
	mux sync.Mutex
}

func (c *Clients) get() map[string]*client {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.client
}

func (c *Clients) getUser (key string) *client {
	c.mux.Lock()
	defer c.mux.Unlock()
	if v, ok := c.client[key]; ok {
		return v
	} else {
		return nil
	}
}

func (c *Clients) addUser(u *client) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.client[u.flag] = u
}

func (c *Clients) delUser(u *client) {
	c.mux.Lock()
	defer c.mux.Unlock()
	delete(c.client, u.flag)
}

// 新建用户
func NewUser(flag, name string) {
	user := &client{name: name, flag: flag}
	clients.addUser(user)
}

// 用户退出
func (c *client) QuitUser() {
	clients.delUser(c)
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
