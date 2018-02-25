package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	//	Server()
	for {
		Client()
		time.Sleep(time.Hour * 3)
	}
}

var channel = make(chan int)

const (
	HOST = "https://www.xxbiquge.com"
	UA   = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36"
	AL   = "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7"
)

type novel struct {
	url,novelName	string
	isDirectory 	bool
	document    	*goquery.Document
}

func Server() {
	fmt.Println("http server start run")
	mux := http.NewServeMux()
	mux.HandleFunc("/user", userHandle)
	http.ListenAndServe(":3000", mux)
}

func userHandle(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("is user"))
}

func Client() {
	directory, err := ioutil.ReadFile("directory.txt")
	if err != nil {
		fmt.Println("error: ", "获取目录文件失败: ", err)
		return
	}
	list := strings.Fields(string(directory))
	for _, val := range list {
		if val == "" {
			continue
		}
		number := strings.Split(val, "--")
		item := novel{url: "/" + number[0] + "/", isDirectory: true, novelName: number[1]}
		item.repeatGet()

	}
	fmt.Println(time.Now(), "处理完成")
}

func (n *novel) repeatGet() {
	for i :=0; ; i++{
		err := n.get()
		if err != nil {
			if i > 5 {
				fmt.Println("error: ", n.url, ": html获取失败->", err)
				break
			}
			time.Sleep(time.Second * 1)
		} else {
			break
		}
	}

}

func (n *novel) get() error {
	req, err := http.NewRequest("GET", HOST+n.url, nil)
	if err != nil {
		fmt.Println("error: ", n.url, ": 链接失败->", err)
		return nil
	}
	req.Header.Set("User-Agent", UA)
	req.Header.Set("Accept-Language", AL)
	client := &http.Client{}
	resp, err := client.Do(req)
	html, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return err
	}
	n.document = html
	if n.isDirectory {
		n.parseHtml()
	} else {
		n.parseContentHtml()
	}
	return nil
}

func (n *novel) parseHtml() {
	list := n.document.Find("#list a")
	listSlice := make([][2]string, len(list.Nodes))

	fileName := strings.Trim(n.url, "/")
	old := getFileContent(fileName + ".txt")

	contentNum := 0
	list.Each(func(i int, content *goquery.Selection) {
		if href, result := content.Attr("href"); true == result {
			listSlice[i] = [2]string{content.Text(), href}
			if false == find(href, &old) {
				content := novel{url: href, isDirectory: false, novelName: n.novelName  + "->" + content.Text()}
				contentNum ++
				go content.repeatGet()
			}
		}
	})
	j, err := json.Marshal(listSlice)
	if err != nil {
		fmt.Println("error :", n.url, ": 转json失败->", err)
	}
	ioutil.WriteFile(fileName+".txt", j, 0666)
	os.Mkdir(fileName, 0755)
	for contentNum > 0  {
		<-channel
		contentNum --
	}
}

func find(val string, array *[][2]string) bool {
	for _, v := range *array {
		if val == v[1] {
			return true
		}
	}
	return false
}

func getFileContent(fileName string) [][2]string {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error: ", fileName, ": 文件加载失败->", err)
		return nil
	}
	c := make([][2]string, 1)
	json.Unmarshal(file, &c)
	return c
}

func (n *novel) parseContentHtml() {
	fileName := strings.TrimPrefix(n.url, "/")
	fileName = strings.TrimSuffix(fileName, ".html")
	content, _ := n.document.Find("#content").Html()
	ioutil.WriteFile(fileName+".txt", []byte(content), 0666)

	fmt.Println(n.novelName)
	channel <- 1
}
