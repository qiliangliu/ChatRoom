package message

const (
	LoginMesType       = "LoginMes"
	LoginResMesType    = "LoginResMes"
	RegisterMesType    = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
)

// Message 封装一个总的消息类
type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容, 注意是被序列化之后的
}

// LoginMes 一个具体的登录消息类
type LoginMes struct {
	UserId   int    `json:"user_id"`
	UserPwd  string `json:"user_pwd"`
	UserName string `json:"user_name"`
}

// LoginResMes 一个具体的登录响应消息类
type LoginResMes struct {
	Code  int    `json:"code"`  //返回的一个登录响应码
	Error string `json:"error"` //错误信息
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`  //返回的一个注册响应码
	Error string `json:"error"` //错误信息
}
