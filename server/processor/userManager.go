package processor

import (
	"fmt"
)

var userMan *UserMan

type UserMan struct {
	onlineUsers map[int]*UserProcessor
}

func init() {
	userMan = &UserMan{
		onlineUsers: make(map[int]*UserProcessor, 1024),
	}
}

func (this *UserMan) AddOnlineUser(up *UserProcessor) {
	this.onlineUsers[up.UserId] = up
}

func (this *UserMan) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

func (this *UserMan) GetAllOnlineUser() map[int]*UserProcessor {
	return this.onlineUsers
}

func (this *UserMan) GetOnlineUserById(userId int) (up *UserProcessor, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d不存在", userId)
		return
	}
	return
}
