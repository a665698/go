package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	"strconv"
)

var upgrader = websocket.Upgrader{}

const TokenName = "CzcyAutoLogin"

type client struct {
	conn *websocket.Conn
	// message chan []byte
	name string
}

var clients map[string]*client
var addUser chan *client
var addMessage chan []byte

func main() {
	fmt.Println("start")
	//fs := http.FileServer(http.Dir("public"))
	//http.Handle("/", fs)
	http.HandleFunc("/", FileHandle)
	http.HandleFunc("/ws", handel)
	go handelChannel()
	http.ListenAndServe(":3000", nil)
}

func FileHandle(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:TokenName,
		Value:strconv.Itoa(int(time.Now().Unix())),
	}
	http.SetCookie(w, &cookie)
	http.ServeFile(w, r, "public/sk.html")
}

func handel(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie(TokenName)
	if err != nil {
		fmt.Println("token不存在: ", err)
		return
	}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrader错误: ", err)
		return
	}
	var clientItem *client
	if val, ok := clients[token.Value]; !ok {
		fmt.Println("用户不存在,创建")
		// clientItem = &client{conn: c, message: make(chan []byte), name: token.Value}
		clientItem = &client{conn: c, name: token.Value}
		addMessage <- []byte("有一个用户加入")
		addUser <- clientItem
	} else {
		fmt.Println("用户存在")
		clientItem = val
	}

	// client := &client{conn: c, message: make(chan []byte), name: token.Value}
	fmt.Println("--61--", len(clients))
	// go clientItem.writeMes()
	go clientItem.readMes()
}

func handelChannel() {
	clients = make(map[string]*client)
	addUser = make(chan *client)
	addMessage = make(chan []byte, 10)
	for {
		select {
		case message := <-addMessage:
			for _, client := range clients {
				// select {
				// case client.message <- message:
				// default:
				// 	fmt.Println("--79--close")
				// 	close(client.message)
				// 	delete(clients, client.name)
				// }
				// client.conn.WriteMessage(1, message)

				client.writeMes(&message)
			}
		case user := <-addUser:
			clients[user.name] = user
		}
	}
}

func (c *client) readMes() {
	for {
		mt, message, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Println("消息获取失败: ", err, "\n消息类型", mt)
			if mt == -1 {
				fmt.Println("--96--close")
				// close(c.message)
				delete(clients, c.name)
				addMessage <- []byte("有一个用户退出")
			}
			break
		}
		fmt.Println("收到消息: ", string(message), " \n消息类型: ", mt)
		addMessage <- message
	}
}

func (c *client) writeMes(message *[]byte) {
	// for {
	// 	select {
	// 	case message := <-c.message:
	// 		c.conn.WriteMessage(1, message)
	// 	}
	// }
	err := c.conn.WriteMessage(1, *message)
	fmt.Println("---119---", err)
}
