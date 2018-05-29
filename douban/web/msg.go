package web

var msgFlags = map[int]string{
	SUCCESS: "成功",
	ERROR:   "错误",
}

func GetMsg(code int) string {
	if msg, ok := msgFlags[code]; ok {
		return msg
	} else {
		return msgFlags[ERROR]
	}
}
