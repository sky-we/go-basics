package main

import "fmt"

func main() {
	// 一、定义数组，注意容量已声明
	intList := [5]int{1, 2, 3, 4, 5}
	intList[0] = 2
	fmt.Println(intList[1:3])

	// 使用 ... 让编译器推断数组长度
	intList2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(intList2)

	// 二、定义切片，切片三要素 指针、长度、容量
	// 1 从数组创建切片
	sliceByArr := intList[2:4]
	fmt.Println(sliceByArr)

	// 2 var变量声明切片, 注意没有声明容量
	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println(slice, len(slice), cap(slice))

	// 3 make初始化切片
	slice2 := make([]int, 6, 12)
	fmt.Println(slice2, len(slice2), cap(slice2))

	newSlice := slice2[1:3]
	fmt.Println(newSlice)

	// 4 基于数组或已有切片创建切片
	newslice2 := slice2[:]
	fmt.Println(slice2, newslice2)

	// 切片特性：切片动态扩容
	// 1）容量随着数组元素的增加而增加
	var sliceExpand []int                // sliceExpand默认值为nil
	sliceExpand = append(sliceExpand, 1) // 4 append扩容后会创建新的切片
	fmt.Println(sliceExpand, len(sliceExpand), cap(sliceExpand))
	sliceExpand = append(sliceExpand, 2)
	fmt.Println(sliceExpand, len(sliceExpand), cap(sliceExpand))
	sliceExpand = append(sliceExpand, 3, 4, 5)
	fmt.Println(sliceExpand, len(sliceExpand), cap(sliceExpand))

	// 2）切片扩容后会导致原数组和新切片引用发生解绑
	// 当原数组触发扩容的时候，原数组将和切片引用发生解绑，Go会将原数组元素复制到新数组中，并返回新的切片引用
	// 原切片引用继续引用原数组，原切片引用和新切片引用不再共享底层数组
	// 这种“解绑”使得扩容后的切片可以灵活地增加容量，而不会影响到其他仍然引用原底层数组的切片
	newSliceExpand := sliceExpand[:]
	fmt.Println(sliceExpand, newSliceExpand)
	sliceExpand[0] = 10
	fmt.Println(sliceExpand, newSliceExpand)
	// 原长度5 容量6 扩容两个元素，cap自动扩容
	sliceExpand = append(sliceExpand, 6, 7)
	// 将原来的元素第1个值修改
	sliceExpand[0] = 20
	//
	fmt.Println(sliceExpand, newSliceExpand)

}
