package main

import (
	"fmt"
)

//用户的账号
var userId int
//用户的密码
var userPwd string

func main() {
	//接受用户的操作信息
	var key int
	//判断是否循环打印菜单
	var loop = true

	for loop {
		fmt.Println("---------------欢迎登陆多人聊天系统---------------")
		fmt.Println("\t\t\t1 登陆聊天室")
		fmt.Println("\t\t\t2 注册用户")
		fmt.Println("\t\t\t3 退出系统")
		fmt.Println("\t\t\t请输入<1~3>")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("\t\t\t登陆聊天室")
			loop = false
		case 2:
			fmt.Println("\t\t\t注册用户")
			loop = false
		case 3:
			fmt.Println("\t\t\t退出系统")
			loop = false
		default:
			fmt.Println("\t\t\t输入有误，请重新输入")

		}

		if key == 1 {
			fmt.Println("\t\t\t请输入用户账号")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("\t\t\t请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			//把登录写在另一个文件里面
			_ = login(userId, userPwd)

		} else {

			fmt.Println("\t\t\t进行其他操作")
		}
	}

}
