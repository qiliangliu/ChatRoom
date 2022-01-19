package main

import (
	"github.com/qiliangliu/ChatRoom/common/message"
	"github.com/qiliangliu/ChatRoom/common/utils"
	"github.com/qiliangliu/ChatRoom/server/processor"
	"net"
	"fmt"
	"io"
)

type Controler struct {
	Conn net.Conn
}

//ServerProcessMes 函数, 根据客户端发送的不同消息类型，决定调用哪个函数来处理
func (this *Controler) ServerProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录函数
		userProcessor := &processor.UserProcessor{
			Conn: this.Conn,
		}
		err = userProcessor.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册的函数
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func (this *Controler) MainControl() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err == io.EOF {
			fmt.Println("客户端退出，服务器也退出")
			return err
		} else if err != nil {
			fmt.Println("readPkg err: ", err)
			return err
		}

		err = this.ServerProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
