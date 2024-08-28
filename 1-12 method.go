package main

/*
首字符大写 代表外部可以调用
方法和struct必须写在一个go文件中，这样可读性强，代表针对这一类的数据，处理的方法有A，B，C...
方法的本质是对func的封装
*/
import (
	"fmt"
	"go-basics/model"
)

type school struct {
	name    string
	addr    string
	student []*model.Student
}

// GetName 首字符大写 代表外部可以调用
func (s *school) GetName() string {
	return s.name

}
func main() {

	s := &school{
		name: "school",
		addr: "幸福小区",
		student: []*model.Student{
			{
				Name: "哈哈",
				Age:  20,
			},
			{
				Name: "嘿嘿",
				Age:  22,
			},
		},
	}
	fmt.Println(s.GetName())
	fmt.Println(s.student)
	for _, student := range s.student {
		fmt.Println(student.GetName())
	}
}

// &（取地址符）：获取一个变量的内存地址，返回一个指向该变量的指针。
// *（解引用符）：通过指针访问其指向的实际值，也可以用来修改指针所指向的数据。
/*
取地址符（&）与结构体
在创建结构体实例时，通常会用 & 取地址符来创建一个指向结构体的指针。这样做的目的是为了避免复制结构体实例，而是直接操作其地址。

解引用符（*）与方法接收器
在定义方法时，如果接收器是指针类型，则使用 * 来表示指针接收器。这意味着方法内部对接收器的任何修改都会影响到原始结构体实例

在结构体中使用 &：通常用于创建指向结构体的指针，以便更高效地操作大型结构体。
在方法接收器中使用 *：用于指针接收器，使方法能够修改接收者的状态。*/
