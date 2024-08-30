package main

import (
	"fmt"
	"sync"
	"time"
)

func ss(ch chan<- int) {
	i := 0
	for {
		time.Sleep(time.Second)
		select {
		case ch <- i:
			i++
		}
	}
}

func work(ch <-chan int, wg *sync.WaitGroup) {
	ticker := time.NewTicker(2 * time.Second)
	defer wg.Done()
	for {
		time.Sleep(time.Millisecond * 500)
		select {
		case data := <-ch:
			fmt.Println(data)
		case <-ticker.C:
			fmt.Println("worker status ok")
		default:
			fmt.Println("队列为空")
		}

	}

}

func main() {
	wg := sync.WaitGroup{}
	// 带缓冲的通道允许你在没有对应接收方的情况下发送数据
	ch := make(chan int)
	wg.Add(1)
	go ss(ch)
	go work(ch, &wg)
	wg.Wait()
}
