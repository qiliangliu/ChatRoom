package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//创建一个全局的userdao，一个服务器只要一个线程池操作对象就可以了, 在mian中被初始化
var MyUserDao *UserDao

func InitUserDao(pool *redis.Pool) {
	MyUserDao = NewUserDao(pool)
}

//Dao -> data access object 数据访问对象
//UserDao 结构体用来完成对User结构特、redis之间所用操作
type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//考虑一下UserDao应该给我提供哪些方法，让我们来完成相应的操作
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("hget", "Users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOT_EXIST
		}
		return
	}

	user = &User{}
	//把res反序列化成实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal error: ", err)
		return
	}
	return
}

//Login 完成登录校验, 是面向数据库中的数据进行校验
//1. 如果用户的id和pwd都正确，则返回一个user实例
//2. 如果用户的id和pwd有错误，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先从UserDao链接的线程池中取出一根链接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	//这时证明用户是存在的，但是我们接下来要验证用户密码是否正确
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
