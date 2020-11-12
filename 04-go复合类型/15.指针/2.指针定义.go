package main

import "fmt"
/*
数据在内存中的地址称为指针(0xc000016060)：如果一个变量存储了一份数据的指针，我们就称它为指针变量。（变量存储指针）
指针变量的值就是某份数据的地址，这样的一份数据可以是数组、字符串、函数.
int *a, *b, *c;  //a、b、c 的类型都是 int*
2.通过指针变量取得数据
int a = 15;
int *p = &a;
printf("%d, %d\n", a, *p);  //两种方式都可以输出a的值
指针变量(p)存储了数据的地址，通过指针变量能够获得该地址上的数据，格式为：*p
*/
func main() {
	var num int = 10
	var p *int
	p = &num
	fmt.Printf("%p\n", p)//0xc000016060
	fmt.Printf("%p\n", &num)
	fmt.Println(*p)
	*p = 30
	fmt.Println("num=", num)
}

