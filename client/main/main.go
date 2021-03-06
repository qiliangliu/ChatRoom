package main

import (
	"fmt"
	"github.com/qiliangliu/ChatRoom/client/processor"
	"os"
)

//用户的账号
var userId int
//用户的密码
var userPwd string
//用户名
var userName string

func main() {
	//接受用户的操作信息
	var key int

	for true {
		fmt.Println("---------------欢迎登陆多人聊天系统---------------")
		fmt.Println("\t\t1 登陆聊天室")
		fmt.Println("\t\t2 注册用户")
		fmt.Println("\t\t3 退出系统")
		fmt.Println("\t\t请输入<1~3>")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("\t\t登陆聊天室")
			fmt.Println("\t\t请输入用户账号")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("\t\t请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			userProcessor := &processor.UserProcessor {
			}
			_ = userProcessor.Login(userId, userPwd)
		case 2:
			fmt.Println("\t\t注册用户")
			fmt.Println("\t\t请输入用户账号")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("\t\t请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("\t\t请输入用户名")
			fmt.Scanf("%s\n", &userName)
			userProcessor := &processor.UserProcessor {
			}
			_ = userProcessor.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("\t\t退出系统")
			os.Exit(0)
		default:
			fmt.Println("\t\t输入有误，请重新输入")
		}
	}
}
