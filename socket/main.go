package socket

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
	name string
}

var clients map[string]*client
var addUser chan *client
var addMessage chan []byte

type Account struct {
	id int
	user string
	password string
}

var accounts = []Account{
	{1, "admin1", "111111"},
	{2, "admin2", "222222"},
	{3, "admin3", "333333"},
	{4, "admin4", "444444"},
	{5, "admin5", "555555"},
	}

func Main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/ws", wsHandel)
	http.HandleFunc("/login", loginHandel)
	go SocketInit()
	http.ListenAndServe(":3000", nil)
}

func loginHandel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write(GetHttpResult(2))
		return
	}
	user := r.PostFormValue("user")
	password := r.PostFormValue("password")
	fmt.Println(user)
	for _, v := range accounts {
		if v.user == user && v.password == password {
			cookie := http.Cookie{
				Name:TokenName,
				Value:strconv.Itoa(int(time.Now().Unix())),
			}
			http.SetCookie(w, &cookie)
			w.Write(GetHttpResult(0))
			return
		}
	}
	w.Write(GetHttpResult(1))
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:TokenName,
		Value:strconv.Itoa(int(time.Now().Unix())),
	}
	http.SetCookie(w, &cookie)
	http.ServeFile(w, r, "public/sk.html")
}

func wsHandel(w http.ResponseWriter, r *http.Request) {
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
		clientItem.readMes()
	} else {
		val.readMes()
	}
}

func SocketInit() {
	clients = make(map[string]*client)
	addUser = make(chan *client)
	addMessage = make(chan []byte, 10)
	go handelChannel()
}

func handelChannel() {
	for {
		select {
		case message := <-addMessage:
			writeMes(&message)
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

func writeMes(message *[]byte) {
	for _, client := range clients {
		err := client.conn.WriteMessage(1, *message)
		fmt.Println("---119---", err)
	}
}
