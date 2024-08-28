package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

func main() {
	// 短变量赋值
	name := "lucky"

	// string转字符数组
	nameCopy := []byte(name)
	nameCopy[1] = 'L'
	fmt.Println(name, string(nameCopy))

	// 默认值为空字符串
	var school string
	fmt.Println(school)

	// 字符串拼接
	var newName string
	newName = name + " " + "day"
	fmt.Println(newName)

	// 字符串比较
	if name == newName {
		fmt.Println("name == newname")
	}
	if name != newName {
		fmt.Println("name == newname")
	}

	// 反引号支持多行字符
	text := `ABCD
  ETF`
	fmt.Println(text)

	// fmt.Sprintf 格式化str并返回
	fmtStr := fmt.Sprintf("My name is %s", "lw")
	fmt.Println(fmtStr)

	// strings.Builder 高效构建str

	var builder strings.Builder
	builder.WriteString("Hello World")
	builder.WriteString("lw")
	builder.WriteRune('世') // 将单个unicode码点写到builder中
	result := builder.String()
	fmt.Println(result)

	// strings.join 切片转字符串
	element := []string{"apple", "banana", "cherry"}
	elementStr := strings.Join(element, ", ")
	fmt.Println(elementStr)

	// strings.Split 字符串转切片
	elementCopy := strings.Split(elementStr, ", ")
	fmt.Println(elementCopy)

	// bytes.Buffer 高效读取二进制数据和字符 构建字符串
	var buffer bytes.Buffer

	buffer.WriteString("Hello")
	buffer.WriteString("world")
	buffer.WriteRune('世') // 将单个unicode码点写到buffer中

	// 拼接二进制数
	var buffer2 bytes.Buffer
	var num int32 = 12345

	// 将num以二进制格式写入到buffer2中
	binary.Write(&buffer2, binary.LittleEndian, num)
	results2 := buffer2.Bytes()
	fmt.Println(string(results2))
}
