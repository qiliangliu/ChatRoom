package processor

import (
	"encoding/json"
	"fmt"
	"github.com/qiliangliu/ChatRoom/common/message"
	"github.com/qiliangliu/ChatRoom/common/utils"
	"net"
	"os"
)

//主要作用
//1. 打印界面
//2. 保持与服务端的链接(通过创建协程)
//3. 接受服务端的推送的消息，并处理

func showmenu() {
	for {
		//接受用户的操作信息
		var key int
		fmt.Println("---------------恭喜xxx登陆聊天系统---------------")
		fmt.Println("\t\t\t1 显示在线列表")
		fmt.Println("\t\t\t2 发送消息")
		fmt.Println("\t\t\t3 信息列表")
		fmt.Println("\t\t\t4 退出系统")
		fmt.Println("\t\t\t请输入<1~4>")
		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("\t\t\t1 显示在线列表")
			outputOnlineUser()
		case 2:
			fmt.Println("\t\t\t2 发送消息")
		case 3:
			fmt.Println("\t\t\t3 信息列表")
		case 4:
			fmt.Println("\t\t\t4 退出系统")
			os.Exit(0)
		default:
			fmt.Println("\t\t\t输入有误，请重新输入")
		}
	}
}

func acceptServerMes(Conn net.Conn) {
	//创建一个Transfer实例，不停的接受服务器推送过来的消息
	tf := &utils.Transfer{
		Conn: Conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		fmt.Println("$$$$$$$$$$$$$$$$$$$$$\n")
		if err != nil {
			fmt.Println("tf.ReadPkg() error: ", err)
			return
		}
		//成功读取到消息，之后应该对消的内容进行相应的处理
		//fmt.Println("mes = %v", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType: //有人上线了
			var notifyUserStatusMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("json.Unmarshal error: ", err)
				return
			}
			updateUserStatus(&notifyUserStatusMes)
			outputOnlineUser()
		default:
			fmt.Println("服务器端返回了未知的消息类型")
		}
	}
}
