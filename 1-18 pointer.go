package main

import "fmt"

type MyError2 struct {
	Code int
	Msg  string
}

// 什么时候使用指针？
//1 需要修改结构体字段：如果需要在方法或函数中修改结构体的字段，应该使用指针接收器或指针参数。

func (e *MyError2) UpdateCode(newCode int) {
	e.Code = newCode
}

//2 避免内存拷贝：如果结构体很大（包含很多字段），传递值类型会复制整个结构体，消耗较多内存。使用指针可以避免这种情况，因为指针只占用少量内存（一般是 8 字节，存储地址）。

func ProcessError(err *MyError2) {
	// 处理错误
}

//3 一致性：如果结构体实现了接口，通常会使用指针接收器，以确保接口的实现一致。

func (e *MyError2) Error() string {
	return e.Msg
}

// 什么时候不使用指针？
// 1 结构体是不可变的：

func UnChange() {
	err := MyError2{Code: 10001, Msg: "Immutable error"}
	fmt.Println(err)

}

//2 结构体很小：

func SmallStruct(code int, msg string) MyError2 {
	return MyError2{Code: code, Msg: msg}
}
