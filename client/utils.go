package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/qiliangliu/ChatRoom/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf[:4])
	if n != 4 || err != nil {
		return
	}

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
