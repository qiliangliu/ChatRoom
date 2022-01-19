package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/qiliangliu/ChatRoom/common/message"
	"github.com/qiliangliu/ChatRoom/common/utils"
	"net"
)

func login(userId int, userPwd string) (err error) {
	//准备链接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}
	//延时关闭链接，不要忘记！！！
	defer conn.Close()

	//链接成功，通过conn来发送消息
	var mes message.Message
	//赋值Type
	mes.Type = message.LoginMesType
	//创建LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//讲loginMes 进行序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err : err")
		return
	}
	//赋值mes.Data
	mes.Data = string(data)
	//讲mes 进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err : err")
		return
	}

	//到这我们先发送数据长度，在发送数据内容
	var pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write len err: ", err)
		return
	}

	fmt.Println("客户端发送消息成功, 发送长度", len(data))
	fmt.Println("发送内容：", string(data))

	//发送消息本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write data err: ", err)
		return
	}

	//这里还需要处理服务器的相应消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err: ", err)
		return
	}
	//将mes反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data)err: ", err)
		return
	}
	//根据响应消息，客户端做出响应的显示
	if loginResMes.Code == 100 {
		fmt.Println("登录成功")
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}
