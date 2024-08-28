package main

import "fmt"

// 有类型常量，不推荐
type newInt int

const Y newInt = 10

// 无类型常量, 代码简洁，推荐

const (
	constName = 10
)
const constName2 = 11

func main() {
	fmt.Println(name)
}
