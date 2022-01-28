package processor

import (
	"encoding/json"
	"fmt"
	"github.com/qiliangliu/ChatRoom/common/message"
	"github.com/qiliangliu/ChatRoom/common/utils"
	"net"
)


type ShortMesProcess struct {
}

func (this *ShortMesProcess) PushGroupMes(mes *message.Message) {
	//取出mes中的ShortMes
	var shortMes message.ShortMes
	err := json.Unmarshal([]byte(mes.Data), &shortMes)
	if err != nil {
		fmt.Println("smsProcessor json.Unmarshal error: ", err)
		return
	}

	//将Mes再次进行序列化, 作为要发送的数据
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("smsProcessor json.Unmarshal error: ", err)
		return
	}

	//遍历onlineUser把消息进行推送到客户端
	fmt.Println("服务器推送ShortMes了￥￥￥￥￥￥￥￥￥￥￥")
	for id, up := range userMan.onlineUsers {
		if id == shortMes.UserId {
			continue
		}
		this.PushMesToOnlineUsers(data, up.Conn)
	}
}

func (this *ShortMesProcess) PushMesToOnlineUsers(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("smsProcessor 转发消息失败")
	}
	return
}
