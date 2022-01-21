package processor

import (
	"fmt"
	"github.com/qiliangliu/ChatRoom/common/message"
)

//客户端要维护的在线用户map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 15)

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
}

//在客户端显示当前在线的用户
func outputOnlineUser() {
	fmt.Println("当前在线用户列表：")
	for id := range onlineUsers {
		fmt.Println("用户id：\t", id)
	}
}