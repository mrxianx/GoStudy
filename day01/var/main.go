package main

import "fmt"

// 批量申明(全局变量) 函数外的每个语句都必须以关键字开始（var,const,func）
var (
	name string
	age  int
	isOk bool
)

func main() {
	var jbbl string
	// 赋值
	name = "张三"
	age = 18
	isOk = true
	jbbl = "局部变量，必须使用，否则报错"

	fmt.Printf("name:%s", name) // %s:占位符 使用name这个变量的值去替换占位符
	fmt.Println("age:", age)
	fmt.Println("isOk:", isOk)
	fmt.Print(jbbl) // Print:在终端中输出打印的内容

	// 声明变量同时赋值
	var addr string = "cd"
	fmt.Println("地址：", addr)

	// 类型推导（根据值判断类型）
	var email = "1782800572@qq.com"
	fmt.Println("邮件：", email)

	// 简短变量声明(仅在函数内部使用，外部不能使用)
	phone := 18482357470
	fmt.Println(phone)

}
