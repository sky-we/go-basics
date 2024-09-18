package main

import (
	"fmt"
	"strconv"
)

/*
使用 strconv.Itoa：
当你只需要将 int 类型的整数转换为字符串时，strconv.Itoa 更加简洁和高效。
如果性能是一个关键因素，特别是在高频次调用时，strconv.Itoa 更合适。

使用 fmt.Sprintf：
当你需要将整数与其他类型的数据（如浮点数、布尔值、字符串等）一起格式化为一个字符串时，fmt.Sprintf 更加灵活和方便。
如果需要更复杂的格式化需求（如填充、对齐、精度控制），fmt.Sprintf 提供了更多的格式化选项
*/
func main() {
	number := 123
	// int ---> str
	str := strconv.Itoa(number)
	fmt.Println("int as a string is:", str)
	// fmt.Sprintf能达到一样的效果
	fmt.Println(fmt.Sprintf("int as a string is: %d", number))

	// int64 ---> xx进制的str
	number1 := int64(789)
	str2 := strconv.FormatInt(number1, 2)
	fmt.Println("Format integer", str2)

	// str ---> int
	str3 := "123"
	intNum, _ := strconv.Atoi(str3)
	fmt.Printf("string as int is %d\n", intNum)
}
