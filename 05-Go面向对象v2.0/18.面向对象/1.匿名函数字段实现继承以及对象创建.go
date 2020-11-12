package main

import "fmt"

//java c #class创建类，而go没有此字段，用结构体创建类
//属性--成员   方法--函数
type Student struct {
	id int
	name string
	age int
}

/*type Student struct {//StudentOne 和 Teacher存在重复的定义部分，因此可以创建一个类，供这个两个类来继承
	id int
	name string
	age int
	score float64
}

type Teacher struct {
	id int
	name string
	age int
	salary float64
}
*/
//创建三个类 Person为父类，StudentOne、TeacherOne为子类，用来继承Person
type Person struct{
	id int
	name string
	age int
}
type StudentOne struct {
	Person        //通过匿名字段实现继承
	score float64
}
type TeacherOne struct {
	Person
	salary float64
}

func main(){
	var stu Student = Student{101,"张三",20}
	fmt.Println(stu)

	//全部初始化
	var stuOne StudentOne = StudentOne{Person{101, "zhangsan", 18}, 100}
	fmt.Println(stuOne)

	//部分初始化
	var stuTwo StudentOne = StudentOne{score:99}
	fmt.Println(stuTwo)

	var stuThree StudentOne = StudentOne{Person:Person{id:103}}
	fmt.Println(stuThree)


}
