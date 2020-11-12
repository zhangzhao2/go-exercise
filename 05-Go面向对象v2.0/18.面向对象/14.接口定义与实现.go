package main

import "fmt"

//定义接口
type Skills interface {
	Running()
	Getname() string
}

type Student1 struct {
	Name string
	Age int
}

//实现接口
func (p Student1) Getname() string{
	fmt.Println("Getname:", p.Name)
	return p.Name
}
func (p Student1) Running() {
	fmt.Println("Running:", p.Name)

}


func main() {
	var skill Skills
	var stu1 Student1
	stu1.Name = "wd"
	stu1.Age = 22
	skill = stu1
	skill.Running()//调用接口
}
