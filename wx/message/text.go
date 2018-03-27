package message

import (
	"wx/config"
	"wx/menu"
	"wx/media"
	"fmt"
	"strconv"
)

type TextMessage struct {
	BaseMessage
	Content string
}

func (m *Message) TextMessageHandle() *TextMessage {
	content := "你输入的是：" + m.Content
	if b, err := strconv.Atoi(m.Content); err == nil {
		if h := GetHandel(b); h != nil {
			content = h()
		}
	}
	//switch m.Content {
	//case "设置菜单":
	//	content = menu.Create()
	//case "删除菜单":
	//	content = menu.Delete()
	//default:
	//	content = "你输入的是：" + m.Content
	//}
	i := m.NewTextMessage(content)
	return i
}

func (m *Message) NewTextMessage (content string) *TextMessage {
	t := &TextMessage{}
	t.Content = content
	t.MsgType = config.MsgTypeText
	m.BaseHandle(&t.BaseMessage)
	return t
}


type Value struct {
	name string
	handle func() string
}

var k []Value

func init() {
	k = append(k, Value{"快捷操作", GetKey})
	k = append(k, Value{"设置菜单", menu.Create})
	k = append(k, Value{"删除菜单", menu.Delete})
	k = append(k, Value{"上传临时素材", media.AddTmpMedia})
	//k[0] =
	//k[1] = Value{"设置菜单", menu.Create}
	//k[2] = Value{"删除菜单", menu.Delete}
	//k[3] = Value{"上传临时素材", media.AddTmpMedia}
	fmt.Println(k)
}

func GetHandel(key int) func() string {
	if h := k[key]; h.handle != nil {
		return h.handle
	} else {
		return nil
	}
}

func GetKey() string {
	fmt.Println(k)
	content := ""
	for key, value := range k  {
		fmt.Println(key)
		content += fmt.Sprintf("%d: %s\n", key, value.name)
	}
	return content
}
