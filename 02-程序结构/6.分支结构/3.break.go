package main

import "fmt"

func main() {
	//1.要求用户输入用户名和密码，只要不是admin、8888888就一直提示用户名，密码错误.请重新输入。
	//
	var userName string
	var userPwd string
	var count int

	for ; ; {
		fmt.Println("请输入用户名:")
		fmt.Scan(&userName)
		fmt.Println("请输入密码:")
		fmt.Scan(&userPwd)

		//2.对输入的用户名和密码进行判断
		if userName == "admin" && userPwd == "888888"{
			fmt.Println("输入的用户名密码正确")
			break
		} else {
			count++
			if count >=3{
				fmt.Println("最多输错3次")
				break
			}
			fmt.Println("输入用户名密码错误，请重新输入")
		}

	}
}