package main

import "fmt"

//定义接口
type Wa interface {
	Running()
	Getname() string
}

type Student2 struct {
	Name string
	Age int
}

type Teacher struct {
	Name string
	Salary int
}
//多态：同一个interface, 不同的类（结构体）进行调用，实现方法不同。就是接口少定义，接口里面的函数那个结构体用，那个结构体自己就定义。
//实现接口
func (p Student2) Getname() string{
	fmt.Println("s-Getname:", p.Name)
	return p.Name
}
func (p Student2) Running() {
	fmt.Println("S-Running:", p.Name)

}

//实现Teacher结构体中的接口方法
func (p Teacher) Getname() string{
	fmt.Println("T_Getname:", p.Name)
	return p.Name
}

func (p Teacher) Running() {
	fmt.Println("T-running:", p.Name)
}

func main() {
	var skill Wa
	var stu1 Student2
	var t1 Teacher
	t1.Name = "wanglaoshi"
	stu1.Name = "litongxue"
	stu1.Age = 22
	skill = stu1
	skill.Running()//调用接口
	skill = t1
	skill.Running()
}
