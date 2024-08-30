package main

import "fmt"

/*
struct 的作用
1 定义数据结构：
struct 是一种聚合数据类型，可以将多个不同类型的字段组合在一起，定义一个数据结构。它通常用于描述某个实体或概念。

2 方法的接收者类型：
struct 可以与方法一起使用，使得这些方法可以在特定的数据类型上操作。

3 分组相关数据：
struct 可以将相关数据字段分组，形成一个逻辑单元，方便数据的组织和传递。
*/

type Request struct {
	method  string
	url     string
	body    []string
	GetBody func(a string, b string) []byte // 函数
}

func main() {
	// 访问结构体
	var req Request
	fmt.Println(req.method, req.url, req.body)

	// 通过字面量直接定义
	s := Request{
		method: "GET",
		url:    "http://www.baidu.com",
		body:   []string{"1", "2", "3"},
	}
	s.GetBody = func(a string, b string) []byte {
		return []byte(a + " " + b)
	}

	// 通过.直接访问
	fmt.Println(s.url)
	result := s.GetBody("1", "2")
	fmt.Println(string(result))

	// 取指针访问
	p := &s

	fmt.Println(p.body)
}
