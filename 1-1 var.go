package main

import "fmt"

/* 全局变量声明 */

// 声明变量
var name string
var age int

// 立刻赋值, 用的比较少
var x = 1
var y = 3.5

// 直接赋值并声明类型
var m int32 = 1
var n float32 = 1.2

// 变量块，将同一类变量合并，增加可读性
var (
	me  int64   = 1
	you float64 = 1.2
)

func main() {
	/* 局部变量声明 */

	// 短变量，立刻赋值，与if语句结合，将变量的作用域最小化
	a := 1
	b := 2
	fmt.Printf("a = %d, b= %d\n", a, b)

	// 延迟赋值
	var c, d int
	c = 1
	d = 2
	fmt.Printf("c = %d, d= %d\n", c, d)

	// 变量块，将同一类变量合并，增加可读性
	var (
		e int     = 1
		f float32 = 1.2
	)
	var (
		g = 1
		h = 1.2
	)
	fmt.Printf("e = %d, f= %f, g= %d,h= %f\n", e, f, g, h)
}
