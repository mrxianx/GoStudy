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

// 插入
const (
	b1 = iota // 0
	b2 = 100  //100
	b3 = iota //2(const 中每新增一行自增加一)
	b4        //3
)

const (
	c1 = iota // 0
	c2 = 100  //100
	c3 = 101  //101
	c4 = iota //3
	c5        //4
)

// 多个常量定义在一行
const (
	e1, e2 = iota + 1, iota + 2 //e1=1,e2=2
	e3, e4 = iota + 1, iota + 2 //e3=2,e4=3
)

// 定义数组集
const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func main() {
	// a = 5 不能覆盖常量的值

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)
	// fmt.Println(f)
	// fmt.Println(g)

	// fmt.Println(a1)
	// fmt.Println(a2)
	// fmt.Println(a3)
	// fmt.Println(a4)

	// fmt.Println(b1)
	// fmt.Println(b2)
	// fmt.Println(b3)
	// fmt.Println(b4)

	// fmt.Println(c1)
	// fmt.Println(c2)
	// fmt.Println(c3)
	// fmt.Println(c4)
	// fmt.Println(c5)

	// fmt.Println(e1)
	// fmt.Println(e2)
	// fmt.Println(e3)
	// fmt.Println(e4)

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
}
