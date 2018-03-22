package menu

import (
	"wx/config"
)


type Button struct {
	Button *[]Button `json:"button,omitempty"`
	SubButton *[]Button `json:"sub_button,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Key string `json:"key,omitempty"`
	Url string `json:"url,omitempty"`
	MediaId string `json:"media_id,omitempty"`
}


func (m *Button) setClickButton(name, key string) {
	m.Button = nil
	m.SubButton = nil
	m.Name = name
	m.Type = config.MenuClick
	m.Key = key
	m.Url = ""
	m.MediaId = ""
}

func (m *Button) setViewButton(name, key string) {
	m.Button = nil
	m.SubButton = nil
	m.Name = name
	m.Type = config.MenuClick
	m.Key = key
	m.Url = ""
	m.MediaId = ""
}


//// 创建自定义菜单
//func Create() {
//	var button1 Menu
//	button1.Type = "click"
//	button1.Name = "测试按钮1"
//	button1.Key = "button1"
//	var subButton1 Menu
//	subButton1.Type = "click"
//	subButton1.Name = "测试子按钮1"
//	subButton1.Key = "sub_button1"
//	var subButton2 Menu
//	subButton2.Type = "click"
//	subButton2.Name = "测试子按钮2"
//	subButton2.Key = "sub_button2"
//	button2 := make(map[string]interface{})
//	button2["name"] = "测试按钮2"
//	button2["sub_button"] = []Menu{subButton1, subButton2}
//	button := make([]interface{}, 2)
//	button[0] = button1
//	button[1] = button2
//	result, err := json.Marshal(map[string]interface{}{"button": button})
//	if err != nil {
//		fmt.Println("menu json make error")
//	}
//	post(CreateMenu + "?access_token=" + accessToken.Token, string(result))
//}
