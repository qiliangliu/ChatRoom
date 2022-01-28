package processor

import (
	"encoding/json"
	"fmt"
	"github.com/qiliangliu/ChatRoom/common/message"
)

//接收到服务器推送的消息之后，我们我们只需要把消息显示出来即可
func outPutGroupMes(mes *message.Message)  {
	var shortMes message.ShortMes
	err := json.Unmarshal([]byte(mes.Data), &shortMes)
	if err != nil {
		fmt.Println("shortMesManager json.Unmarshal error: ", err)
		return
	}

	//显示消息
	info := fmt.Sprintf("用户Id：%d，对大家说：%s\n", shortMes.UserId, shortMes.Content)
	fmt.Println(info)
}
