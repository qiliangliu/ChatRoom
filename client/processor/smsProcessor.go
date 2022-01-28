package processor

import (
	"encoding/json"
	"fmt"
	"github.com/qiliangliu/ChatRoom/common/message"
	"github.com/qiliangliu/ChatRoom/common/utils"
)

type ShortMesProcessor struct {
}

//把要发送的消息进行序列化并发送
func (this *ShortMesProcessor) SendGroupMes(content string) (err error) {
	//1. 创建一个mes
	var mes message.Message
	mes.Type = message.ShortMesType
	//2. 创建一个shortMes
	var shortMes message.ShortMes
	shortMes.Content = content
	shortMes.UserId = CurUser.UserId
	shortMes.UserStatus = CurUser.UserStatus
	//3. 序列化shortMes
	data, err := json.Marshal(shortMes)
	if err != nil {
		fmt.Println("smsProcessor json.Marshal error: ", err)
		return
	}

	mes.Data = string(data)
	//4. 序列化mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("smsProcessor json.Marshal error: ", err)
		return
	}

	//5. 将序列化后的mes发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("smsProcessor tf.WritePkg error: ", err)
		return
	}
	return
}
