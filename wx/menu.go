package wx

import (
	"encoding/json"
	"fmt"
)

type MenuMessage struct {
	BaseResponse
	Event string
	EventKey string
}

type Menu struct {
	Type string `json:"type"`
	Name string	`json:"name"`
	Key string `json:"key"`
}

// 创建自定义菜单
func MenuCreate() {
	var button1 Menu
	button1.Type = "click"
	button1.Name = "测试按钮1"
	button1.Key = "button1"
	var subButton1 Menu
	subButton1.Type = "click"
	subButton1.Name = "测试子按钮1"
	subButton1.Key = "sub_button1"
	var subButton2 Menu
	subButton2.Type = "click"
	subButton2.Name = "测试子按钮2"
	subButton2.Key = "sub_button2"
	button2 := make(map[string]interface{})
	button2["name"] = "测试按钮2"
	button2["sub_button"] = []Menu{subButton1, subButton2}
	button := make([]interface{}, 2)
	button[0] = button1
	button[1] = button2
	result, err := json.Marshal(map[string]interface{}{"button": button})
	if err != nil {
		fmt.Println("menu json make error")
	}
	post(CreateMenu + "?access_token=" + accessToken.Token, string(result))
}

func (m *MenuMessage) Handle() {

}
