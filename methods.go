package main

type school struct {
	name string
	addr string
}

func (s *school) GetName() string {
	return s.name

}
func main() {

}

// &（取地址符）：获取一个变量的内存地址，返回一个指向该变量的指针。
// *（解引用符）：通过指针访问其指向的实际值，也可以用来修改指针所指向的数据。
