package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer func() {
			// recover与defer配合使用恢复panic并打印堆栈
			if err := recover(); err != nil {
				fmt.Println("err = ", err)
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				fmt.Println(err, string(buf))
			}
		}()
		// 会触发panic
		var names []string
		fmt.Println(names[0])
	}()
	time.Sleep(5 * time.Second)
}
