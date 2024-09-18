package main

/*
首字符大写 代表外部可以调用
方法和struct最好写在一个go文件中，这样可读性强，代表针对这一类的数据，处理的方法有A，B，C...
方法的本质是对func的封装
*/
import (
	"fmt"
	"go-basics/model"
)

type School struct {
	name    string
	addr    string
	student []*model.Student
}

// GetName 首字符大写 代表外部可以调用
func (s *School) GetName() string {
	return s.name

}
func main() {

	s := &School{
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
