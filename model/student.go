package model

// 什么时候定义值类型，什么时候定义引用类型
/*
	值接收器
	1 当方法只需要读取接收者的状态，不希望方法修改接收者，
	2 接收者较小且复制开销不大
	指针接收器
	1 当方法需要修改接收者的状态，
	2 接收者较大或复制开销较高，避免不必要的对象复制
*/

type Student struct {
	Age  int
	Name string
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) SetName(name string) {
	s.Name = name
}
