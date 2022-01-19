package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/qiliangliu/ChatRoom/common/message"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf[:4])
	if n != 4 || err != nil {
		return
	}
	fmt.Println("成功读取到buf = ", buf[:4])
	//把发送的数据长度转换成int32
	var pkgLen = binary.BigEndian.Uint32(buf[:4])
	//把发送的数据读入到buf中去
	n, err = conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	//成功读取数据，但是我们要把数据反序列化后得到message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)		//注意这里要填&mes的地址
	if err != nil {
		fmt.Println("json.Unmarshal err: ", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(byte) fail", err)
		return
	}
	//发送数据本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}
	return
}

//serverProcessLogin 专门处理登录请求函数
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
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
	} else  {
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
	err = writePkg(conn, data)
	return
}

//ServerProcessMes 函数, 根据客户端发送的不同消息类型，决定调用哪个函数来处理
func ServerProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录函数
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
		//处理注册的函数
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}


func process(conn net.Conn) {
	defer conn.Close()
	//循环读客户端发送的信息
	fmt.Println("读取客户端发送的消息...")
	for {
		mes, err := readPkg(conn)
		if err == io.EOF {
			fmt.Println("客户端退出，服务器也退出")
			return
		} else if err != nil {
			fmt.Println("readPkg err: ", err)
		}

		err = ServerProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
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