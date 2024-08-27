package main

import "fmt"

func main() {
	// for循环体
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 7 {
			break
		} else {
			fmt.Println(i)
			continue
		}
	}

	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// 常用与定时任务
	k := 0
	for {
		k++
		fmt.Println(k)
		if k == 10 {
			break
		}
	}

	//for range
	names := []string{"Alice", "Bob", "Charlie"}
	for idx, name := range names {
		fmt.Println(idx, name)
	}

	strNames := "lw"
	for _, v := range strNames {
		fmt.Println(v)
	}

	// for range存在一些问题

	data := []string{"a", "b", "c"}
	// 1）for k v， kv每次遍历时都是用的同一个引用，只创建一次，而不是创建新引用
	var ptrs []*string
	for _, v := range data {
		// value := v 创建新的变量来存储值（值复制）即可解决输出ccc的问题
		ptrs = append(ptrs, &v)
	}

	for _, p := range ptrs {
		fmt.Println(*p) // 输出: c c c
	}

	// 2）for range 遍历映射（map）时，遍历顺序是不确定的
	// Go 中的映射在设计上不保证有序遍历。这是为了提供更好的性能和均匀的哈希分布

	myMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	for key, value := range myMap {
		fmt.Println(key, value) // 输出顺序是不确定的
	}
	// 调用100次
	testRangeMap()

}

func testRangeMap() {

	for i := 0; i < 100; i++ {
		myMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		for key, value := range myMap {
			fmt.Println(key, value) // 输出顺序是不确定的

		}
		fmt.Println("End")
	}

}
