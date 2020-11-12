package main

import "fmt"
/*
map是一种无序的键值对集合。可以通过key值来检索数据，map也是一种集合，所以我们可以像爹地啊数组和切片一样迭代他
map是无序的，我们无法决定它的返回顺序，这是因为Map是通过hash实现的。
*/
func main() {
	var m map[int]string = map[int]string{1:"张三", 2:"李四"}
	DeleteMap(m)
	PrintMap(m)
}

func PrintMap(m map[int]string)  {
	for key, value := range m{
		fmt.Println(key)
		fmt.Println(value)
	}
}

func DeleteMap(m map[int]string)  {
	delete(m, 2)

}
