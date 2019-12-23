// 常量：
package main

import "fmt"

// 定义常量
const a = 17

// 批量定义
const ( // ***批量定义常量时，如果常量为赋值，则该常量为上一个常量的值
	b = 101 //101
	c = 100 //100
	d       //100
)

// *****iota 在const关键字出现时，将被重置为0，const中每新增一行常量声明将是iota自增（iota可以理解为数组中的索引）
const (
	e = iota //0
	f        //1
	g        //2
)

// 面试题
const (
	a1 = iota //0
	a2        //1
	a3        //2
	_         //3
	a4        //4
)

func main() {
	// a = 5 不能覆盖常量的值
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
}
