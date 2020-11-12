package main

import (
	"fmt"
	"strings"
)

func main() {
	//Contains 判断一个字符串是否在另一个字符串中
	var str string = "hellogo"
	b := strings.Contains(str, "goo")
	fmt.Println(b)

	//Join 以"|"对切片中的数据进行分隔
	s := []string{"abc", "hello", "world"}
	str1 := strings.Join(s, "|")
	fmt.Println(str1)

	//Index 判断"Hello"在str中出现的位置，注意从0开始计算。
	str2 := "abcHello"
	n := strings.Index(str2, "Hello")
	fmt.Println(n)

	//Replace 用新的字符串替换旧的字符串，第四个参数表示替换的次数，如果小于0，表示全部替换
	str3 := "Hello World"
	fmt.Println(strings.Replace(str3, "l", "w", -1))

	//Repeat 表示字符串"go"重复的次数
	str4 := strings.Repeat("go", 3)
	fmt.Println(str4)

	//Split 用"@"切割字符串返回切片
	str5 := "hello@world@itcast"
	str5s := strings.Split(str5, "@")
	fmt.Println(str5s)//[hello world itcast]

}
