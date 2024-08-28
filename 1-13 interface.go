package main

import "fmt"

/*interface嵌入到interface中：
接口扩展：通过组合接口，可以创建一个更强大、更灵活的接口，而不需要修改原有的接口定义。
接口复用：复用多个接口的方法，使得新接口自动具备这些方法的行为。
多态性：组合接口也可以利用多态特性，将其传递给接受其组成部分的函数或方法。*/

type Shop interface {
	Manage
	GetShopName() string
}

type Manage interface {
	Offer(s *Staff)
	FindStaffByName(name string) *Staff
}

type Skill interface {
	Greet() string
	account() int
}

type BasicSkill struct {
	SkillName string
}

func (b *BasicSkill) Greet() string {
	return "Hello! Welcome"
}

func (b *BasicSkill) account() int {
	return 100
}

func (b *BasicSkill) GetSkillName() string {
	return b.SkillName
}

type FoodShop struct {
	Name   string
	Price  int
	Staffs []*Staff
}

// struct中嵌入interface
//代码复用：结构体通过嵌入接口，可以复用接口的方法，避免显式实现。
//匿名字段：嵌入的接口可以作为匿名字段，结构体自动具备该接口的方法。
//多态性：结构体可以使用接口的多态特性，可以将结构体实例传递给任何接受该接口的函数或方法。

type Staff struct {
	Name   string
	salary int
	Skill  // 2 嵌入skill interface到struct中
}

func (f *FoodShop) GetShopName() string {
	return f.Name
}

func (f *FoodShop) Offer(s *Staff) {
	f.Staffs = append(f.Staffs, s)
}

func (f *FoodShop) FindStaffByName(name string) *Staff {
	for _, staff := range f.Staffs {
		if staff.Name == name {
			return staff
		}
	}
	return nil
}

func main() {
	var shop Shop = &FoodShop{
		Name:  "四川豌杂面",
		Price: 12,
	}
	fmt.Println(shop.GetShopName())

	basicSkill := &BasicSkill{}
	// 创建员工
	staff := &Staff{
		Name:   "lw",
		salary: 1000,
		Skill:  basicSkill,
	}
	shop.Offer(staff)
	s := shop.FindStaffByName("lw")
	fmt.Println(s)
	fmt.Println(s.Greet())
	fmt.Println(s.account())

}
