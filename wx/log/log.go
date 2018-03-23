package log

import (
	"fmt"
	"time"
)

func MessageLog(content string)  {
	fmt.Println(fmt.Sprint(time.Now(), content))
}
