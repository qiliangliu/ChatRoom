package model

import (
	"github.com/qiliangliu/ChatRoom/common/message"
	"net"
)

//CurUser 因为在客户端，我们有很多地方用到，所以我们将其声明成一个全局变量, 在userManager.go 中
type CurUser struct {
	Conn net.Conn
	message.User
}
