package socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"crypto/md5"
)

var upgrader = websocket.Upgrader{}

const TokenName = "token"

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
	http.ListenAndServe(":3000", nil)
}

func loginHandel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		AjaxReturn(w, 2)
		return
	}
	user := r.PostFormValue("user")
	password := r.PostFormValue("password")
	for _, v := range accounts {
		if v.user == user && v.password == password {
			salt := md5.New()
			salt.Write([]byte(user))
			readyRandom := fmt.Sprintf("%x", salt.Sum(nil))
			UserAdd(readyRandom, user)
			cookie := http.Cookie{
				Name:TokenName,
				Value:readyRandom,
			}
			http.SetCookie(w, &cookie)
			AjaxReturn(w, 0)
			return
		}
	}
	AjaxReturn(w, 1)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie(TokenName)
	if err != nil {
		http.Redirect(w, r, "/public/login.html", http.StatusFound)
		return
	}
	if  clients.getUser(token.Value) == nil {
		fmt.Println("用户不存在")
		http.Redirect(w, r, "/public/login.html", http.StatusFound)
	} else {
		http.ServeFile(w, r, "public/index.html")
	}
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
	if val := clients.getUser(token.Value); val == nil {
		fmt.Println("用户不存在")
	} else {
		if val.conn == nil {
			val.conn = c
		}
		val.readMes()
	}
}

func UserAdd(flag, name string) {
	sendSysMessage(name + " 加入")
	NewUser(flag, name)
}
