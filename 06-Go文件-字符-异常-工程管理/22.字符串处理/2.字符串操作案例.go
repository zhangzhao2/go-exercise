package main

import (
		"fmt"
		"strings"
)
func main() {
	//任务一：输入2018-12-31 输出：2018年12月31日
	fmt.Println("请输入：年-月-日")
	var str string
	fmt.Scan(&str)
	//按照"—"分隔字符
	s := strings.Split(str,"-")
	//输出指定的格式：
	fmt.Println(s[0]+"年"+s[1]+"月"+s[2]+"日")

	//任务二：判断用户输入的内容有没有"傻逼"，有替换为"**"
	fmt.Println("请输入第一句话：")
	var str2 string
	fmt.Scan(&str2)
	if strings.Contains(str2, "傻逼"){
		str2 = strings.Replace(str2, "傻逼", "**", -1)
	}
	fmt.Println(str2)
}
