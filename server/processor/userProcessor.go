package processor

import (
	"encoding/json"
	"fmt"
	"github.com/qiliangliu/ChatRoom/common/message"
	"github.com/qiliangliu/ChatRoom/common/utils"
	"github.com/qiliangliu/ChatRoom/server/model"
	"net"
)

type UserProcessor struct {
	Conn net.Conn
	UserId int
}

//serverProcessLogin 专门处理登录请求函数
func (this *UserProcessor) ServerProcessLogin(mes *message.Message) (err error) {
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
	fmt.Println("账号：", loginMes.UserId)
	fmt.Println("密码：", loginMes.UserPwd)

	//使用model.MyUserDao到redis去验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOT_EXIST {
			loginResMes.Code = 200 //用户不存在
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_EXIST {
			loginResMes.Code = 300 //用户已经存在
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 400 //密码错误
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 500 //未知错误
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 100
		fmt.Println(user, "登录成功")
		//把当前用户加到OnlineUser中去
		this.UserId = loginMes.UserId
		userMan.onlineUsers[this.UserId] = this
		//遍历OnlineUsers把在的人放到UserList中去
		for id, _ := range userMan.onlineUsers {
			loginResMes.UserList = append(loginResMes.UserList, id)
		}
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
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("server tf.WritePkg error: ", err)
	}

	fmt.Println("服务器推送用户上线消息")
	this.NotifyOtherOnlineUser(this.UserId)

	return
}

//serverProcessRegister 专门处理登录请求函数
func (this *UserProcessor) ServerProcessRegister(mes *message.Message) (err error) {
	//1. 先从mes中取出mes.Data，然后反序列求出RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err: ", err)
		return
	}
	//2. 先声明resMes, 用来返回相应消息
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	//2. 再声明一个registerResMes, 用在做间接封装用
	var registerResMes message.RegisterResMes
	fmt.Println("账号：", registerMes.User.UserId)
	fmt.Println("密码：", registerMes.User.UserPwd)
	fmt.Println("姓名：", registerMes.User.UserName)

	//使用model.MyUserDao到redis去验证
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_NOT_EXIST {
			registerResMes.Code = 200 //用户不存在
			registerResMes.Error = err.Error()
		} else if err == model.ERROR_USER_EXIST {
			registerResMes.Code = 300 //用户已经存在
			registerResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			registerResMes.Code = 400 //密码错误
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 500 //未知错误
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 110
		fmt.Println("注册成功, 请返回登录")
	}

	//4. registerResMes 序列化，然后服务器发送响应消息
	data, err := json.Marshal(registerResMes)
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
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//userId 用户通知其他用户，他上线了
func (this *UserProcessor) NotifyOtherOnlineUser(userId int) {
	for id, up :=  range userMan.onlineUsers {
		if id == userId {
			continue
		}
		//开始写一个通知方法
		var mes message.Message
		mes.Type = message.NotifyUserStatusMesType
		var notifyUserStatus message.NotifyUserStatusMes
		notifyUserStatus.UserId = userId
		notifyUserStatus.Status = message.UserOnline
		data, err := json.Marshal(notifyUserStatus)
		if err != nil {
			fmt.Println("json.Marshal error: ", err)
			return
		}
		mes.Data = string(data)
		data, err = json.Marshal(mes)
		if err != nil {
			fmt.Println("json.Marshal error: ", err)
			return
		}
		up.NotifyMeOnline(userId, data)
	}
}

func (this *UserProcessor) NotifyMeOnline(userId int, data []byte) {
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	fmt.Println("打印：链接conn：", tf.Conn)
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline error: ", err)
		return
	}
}