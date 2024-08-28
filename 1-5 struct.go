package main

import "fmt"

// Request 定义结构体
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
