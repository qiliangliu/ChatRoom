package processor

import (
	"encoding/json"
	"fmt"
	"github.com/qiliangliu/ChatRoom/common/message"
	"github.com/qiliangliu/ChatRoom/common/utils"
	"net"
)

type UserProcessor struct {
	Conn net.Conn
}

//serverProcessLogin 专门处理登录请求函数
func (this *UserProcessor) ServerProcessLogin(mes *message.Message) (err error) {
	//1. 先从mes中取出mes.Data，然后反序列求出LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err: ", err)
		return
	}
	//2. 先声明resMes, 用来返回相应消息
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//2. 再声明一个loginResMes, 用在做间接封装用
	var loginResMes message.LoginResMes
	//3. 先把用户账号和密码写死，方便写代码
	fmt.Println("账号：", loginMes.UserId)
	fmt.Println("密码：", loginMes.UserPwd)
	if loginMes.UserId == 1 && loginMes.UserPwd == "1" {
		loginResMes.Code = 100
	} else {
		loginResMes.Code = 200
		loginResMes.Error = "该用户不存在，请注册再使用"
	}

	//4. longinResMes 序列化，然后服务器发送响应消息
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}
	resMes.Data = string(data)

	//5. 对resMes进行序列化，并发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}
	//6. 发送服务器消息，给客户端
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return
}