package main

import (
	"fmt"
	"net"
	"time"
	"github.com/qiliangliu/ChatRoom/server/model"
)

func process(conn net.Conn) {
	defer conn.Close()
	//循环读客户端发送的信息
	fmt.Println("读取客户端发送的消息...")

	mainControl := &Controler{
		Conn : conn,
	}
	_ = mainControl.MainControl()
}

func init() {
	//赋值pool
	initPool("localhost:6379", 16, 0, 300 * time.Second)
	//利用pool给MyUserDao赋值
	model.InitUserDao(pool)
}

func main() {
	//服务器开始监听
	fmt.Println("服务器在8889端口进行监听...")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}
	//监听成功，就等待客户端来链接服务器
	for {
		fmt.Println("等待客户端进行链接...")
		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("listen.Accept err: ", err)
			return
		}
		//一旦链接成功，则启动一个协程和客户端保持通信
		go process(conn)
	}

}