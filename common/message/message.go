package common

// Message 封装一个总的消息类
type Message struct {
	Type string			//消息类型
	Date string			//消息内容, 注意是被序列化之后的
}

// LoginMes 一个具体的登录消息类
type LoginMes struct {
	UserId int
	UserPwd string
	UserName string
}

// LoginResMes 一个具体的登录响应消息类
type LoginResMes struct {
	Code int		//返回的一个响应码
	Error string	//错误信息
}




