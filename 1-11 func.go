package main

import (
	"fmt"
	"io"
	"os"
)

// 服务注册
// 数据库连接
// init函数最先执行
func init() {
	fmt.Println("init func")

}

func main() {
	// main作为函数入口
	fmt.Println("main func")
	data, err := readFile("defer_release.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}

}

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("close file err %s\n", closeErr)
		}
	}()
	data, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}
	return string(data), nil
}
