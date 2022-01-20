package model

//定义一个用户结构体, 是服务器与redis之间数据交流的一个元素
type User struct {
	//注意为了序列化和反序列化，
	//我们必选保证用用户信息json字符串的key和我们的字段名对应的tag是一致的
	UserId int	`json:"user_id"`
	UserPwd string `json:"user_pwd"`
	UserName string `json:"user_name"`
}