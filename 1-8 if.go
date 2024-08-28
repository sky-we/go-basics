package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := rand.Int()
	fmt.Println(num)
	if num%2 == 0 {
		fmt.Println("even")

	} else if num%3 == 0 {
		fmt.Println("odd")

	} else {
		fmt.Println("error")
	}

	// if与字面量结合，缩小变量的作用域,变量作用域只在if语句中
	if num2 := rand.Int(); num2%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// 等同于：
	{
		num2 := rand.Int()
		if num2%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
