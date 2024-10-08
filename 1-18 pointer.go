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

// &（取地址符）：获取一个变量的内存地址，返回一个指向该变量的指针。
// *（解引用符）：通过指针访问其指向的实际值，也可以用来修改指针所指向的数据。
/*
取地址符（&）
1 获取变量的内存地址，返回一个变量指针
2 函数调用时传递变量、结构体、slice、map的地址，以便方便的修改和传递，避免复制整个对象
3 用于创建struct绑定的方法中的指针接收器，这样方法中操作的就是结构体的指针

解引用符（*）
1 用于声明变量是指针类型的，变量声明，方法返回值（数组指针，切片指针，int指针）声明等
2 用于将指针类型变量的值取出，解引用
3 作为函数、方法的参数，声明参数是一个指针接收器，避免直接复制大对象
4 解引用链，**p2 解引用两级指针

Python: 引用是抽象的，用于描述变量和对象的关系，开发者无法直接操作内存地址。
引用不需要显式的指针操作，变量是对对象的引用，所有对象操作都是通过引用间接完成的。
开发者无需管理内存，由python自己管理内存
（python中的id函数以整数形式对内存地址进行了一个抽象表示，而不是实际的内存地址）。
引用的原理是通过Cpython中的PyObject对象来实现，对所有的python对象操作本质是对PyObject结构体中的指针进行操作内存地址

Go: 指针是具体的，开发者可以直接操作内存地址和指针值。指针提供了对内存的直接控制
go也有垃圾回收机制，但允许显式的指针操作，提供了更高的控制内存和性能优化的能力
*/
