package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/qiliangliu/ChatRoom/common/message"
	"net"
)

type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn
	Buf  [2048]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	n, err := this.Conn.Read(this.Buf[:4])
	if n != 4 || err != nil {
		return
	}

	//把发送的数据长度转换成int32
	var pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	//把发送的数据读入到buf中去
	n, err = this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	//成功读取数据，但是我们要把数据反序列化后得到message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes) //注意这里要填&mes的地址
	if err != nil {
		fmt.Println("json.Unmarshal err: ", err)
		return
	}
	return
}

func (this Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(byte) fail", err)
		return
	}
	//发送数据本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}
	return
}
