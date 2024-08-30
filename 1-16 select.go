package main

import (
	"fmt"
)

func s(ch chan<- int) {

	for i := 0; i < 5; i++ {
		select {
		case ch <- i:
			fmt.Println("send ok", i)
		default: // 解决死锁问题
			fmt.Println("channel通道空间不足")
		}
	}
	close(ch)

}

func main() {
	ch := make(chan int)
	s(ch)
}
