package main

import (
	"fmt"
)

type Person struct {
	userName string
	addressPhone map[string]string //key：表示电话类型，value:电话
}

var personList = make([]Person, 0)

func addPerson() {
	//1:定义结构体表示练习疼的信息
	//2:定义切片保存多个人的联系信息。
	//3:向切片中保存数据。
	//3.1 添加姓名

	var name string
	var address string
	var phone string
	var exit string
	var addressPhone = make(map[string]string)

	fmt.Println("请输入姓名")
	fmt.Scan(&name)

	for ; ; {
		//3.2 保存电话类型
		fmt.Println("请输入电话类型")
		fmt.Scan(&address)
		//3.3 保存电话号码
		fmt.Println("请输入电话号码")
		fmt.Scan(&phone)
		//将电话以及电话类型存储到addressPhone中。
		addressPhone[address] = phone
		fmt.Println("如果结束电话的录入，请按Q")
		fmt.Scan(&exit)
		if exit == "Q" {
			break;
		}else {
			continue
		}
	}//将联系人存储到切片中
	personList = append(personList, Person{userName:name, addressPhone:addressPhone})
	showPersonList()//调用函数展示联系人的信息
}

func showPersonList() {
	//1:判断一下切片中是否有数据
	if len(personList) == 0{
		fmt.Println("暂时没有联系人信息")
	} else {
		//2：可以通过循环的方式打印切片中的数据。
		for _, value := range personList{
			fmt.Println("姓名：", value.userName)
			for k, v := range value.addressPhone{
				fmt.Println("电话类型：", k)
				fmt.Println("电话号码：", v)
			}
		}
	}
}

func removePerson()  {
	//1.请输入要删除人的姓名
	var name string
	var index int = -1 //记录找到的联系人信息在切片的下标
	fmt.Println("请输入删除人姓名")
	fmt.Scan(&name)
	//2:根据输入的联系热姓名，查找对应的联系信息
	for k, value := range personList{
		if value.userName == name{
			index = k
			break
		}
	}

	//3:打印输出结果
	if index != -1{
		personList = append(personList[0:index], personList[index+1:]...)//append 函数如果第二个参数为切片，后面要跟三个点
	}
	showPersonList()
}

func findPerson() *Person{
	fmt.Println("请输入查询人姓名")
	var name string
	fmt.Scan(&name)
	var index int = -1

	for k, value := range personList{
		if value.userName == name {
			index = k
			fmt.Println("姓名：", name)
			for k1, value1 := range value.addressPhone{
				fmt.Println("电话类型：", k1)
				fmt.Println("电话号码：", value1)
			}
		}
	}
	if index != -1{
		fmt.Println("没有找到联系人")
		return nil
	}else {
		return &personList[index]
	}

}

func editPerson()  {
	fmt.Println("请输入编辑人姓名")
}

func main() {
	for ; ; {
		scanNum()
	}
}

func scanNum() {
	fmt.Println("添加联系人信息，请按1")
	fmt.Println("删除联系人信息，请按2")
	fmt.Println("查询联系人信息，请按3")
	fmt.Println("编辑联系人信息，请按4")
	//fmt.Println()

	//2：对用户输入的数字进行判断。
	var num int//保存用户输入的数字
	fmt.Scan(&num)
	switchType(num)
}

func switchType(n int) {
	switch n {
	case 1:
		//添加联系人的操作
		addPerson()
	case 2:
		//删除联系人的操作
		removePerson()
	case 3:
		//查询联系人的操作
		findPerson()
	case 4:
		//编辑联系人的操作
		editPerson()
	}
}