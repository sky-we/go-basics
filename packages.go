package main

import (
	"fmt"
	customFmt "go-basics/fmt" // 通过别名处理外部包的重名问题
)

func main() {

	// 导入自己工程下的包之前得初始化go-basics模块 go mod init go-basics
	// Go 能够正确识别和解析你项目中的包，并确保包路径在整个项目中是唯一且清晰的。
	// 通过模块路径，Go 知道如何从项目中正确导入和使用自定义包。
	customFmt.TooLarge(3)
	fmt.Println("Hello World")
}
