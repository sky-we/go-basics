package main

import "fmt"

func main() {
	// 字面量创建
	names := map[int]string{
		1: "ming",
		2: "li",
		3: "liu",
	}
	fmt.Println(names)

	// make创建(已分配内存）
	ages := make(map[int]int)
	ages[1] = 1
	ages[2] = 2
	ages[3] = 3
	fmt.Println(ages)
	delete(ages, 1)   // 删除
	names[1] = "haha" // 修改
	fmt.Println(ages)

	for k, v := range names {
		fmt.Println(k, v)
	}

	// 1） map需要make关键字分配内存以后才可以添加值 不然为map只为nil
	var testMap map[int]string
	// testMap[0] = "xiaoming"
	fmt.Println(testMap)

	// 2） map也可以指定容量,并且也有动态扩容的机制
	ageMap := make(map[int]int, 10)
	fmt.Println(len(ageMap))
}
