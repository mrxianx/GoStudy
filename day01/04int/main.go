package main

import "fmt"

func main() {

	// 类型推导
	var i = 101
	fmt.Printf("%d\n", i)
	fmt.Printf("%b\n", i) // 把十进制转换成二进制
	fmt.Printf("%o\n", i) // 把十进制转换成八进制
	fmt.Printf("%x\n", i) // 把十进制转换成十六进制

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	// 十六进制
	i3 := 0x1234566
	fmt.Printf("%d\n", i3)

	// 查看变量类型
	fmt.Printf("%T\n", i3)

	i4 := int8(9)
	fmt.Printf("%T\n", i4)

}
