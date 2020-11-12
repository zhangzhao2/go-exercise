package main

import "fmt"
//map 类似key value 性质
func main() {
	var m map[int]string = map[int]string {1:"张三", 2:"李四", 3:"王五", 4:"李六"}
	m1 := map[int]string {1:"张三", 2:"李四", 3:"王五", 4:"李六"}
	m2 := make(map[int]string, 10)
	m2[1] = "张三"
	m2[2] = "李四"
	m2[3] = "李四"
	fmt.Println("m=", m)
	fmt.Println("m1=", m1)
	fmt.Println("m2=", m2)


}
