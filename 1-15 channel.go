package main

/*
不带缓冲channel
发送阻塞：当一个 goroutine 试图向不带缓冲的通道发送数据时，它会被阻塞，直到有另一个 goroutine 从通道中接收数据为止。
接收阻塞：当一个 goroutine 试图从不带缓冲的通道接收数据时，它会被阻塞，直到有另一个 goroutine 向通道发送数据为止。

带缓冲channel
发送非阻塞：当一个 goroutine 向带缓冲的通道发送数据时，如果缓冲区未满，发送操作不会被阻塞。只有当缓冲区满时，发送操作才会阻塞，直到缓冲区中有空闲空间。
接收非阻塞：当一个 goroutine 从带缓冲的通道接收数据时，如果缓冲区非空，接收操作不会被阻塞。只有当缓冲区为空时，接收操作才会阻塞，直到有新的数据发送到通道中。

不带缓冲通道 vs 带缓冲通道的区别
同步 vs 异步：

不带缓冲通道：严格同步操作。发送和接收操作必须成对出现，才能成功进行。发送方和接收方必须等待对方。
带缓冲通道：允许异步操作。发送方可以在接收方处理之前完成发送，只要缓冲区未满。同样，接收方可以在发送方发送更多数据之前处理已有的数据，只要缓冲区非空。
典型用途：

不带缓冲通道：适用于需要精确控制并发的场景，如要求 goroutine 严格按照顺序执行。
带缓冲通道：适用于允许某些程度的并发异步处理的场景，比如生产者-消费者模式，能平滑过渡高峰期的数据处理。

总结
不带缓冲通道：提供的是同步通信，发送和接收操作必须同时进行，适合需要严格顺序控制的场景。
带缓冲通道：允许异步处理，能在一定程度上解耦发送方和接收方的执行顺序，适合生产者-消费者模型等场景。
*/
import (
	"fmt"
	"sync"
	"time"
)

func send(ch chan<- int) {

	for i := 0; i < 5; i++ {
		fmt.Printf("send data %d,\n", i)
		ch <- i
	}
	close(ch)

}

func receive(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		time.Sleep(1 * time.Second)
		fmt.Printf("receive data %d,\n", data)
	}
}

func main() {
	wg := sync.WaitGroup{}
	// 带缓冲的通道允许你在没有对应接收方的情况下发送数据
	ch := make(chan int)
	wg.Add(1)
	go send(ch)
	time.Sleep(time.Second * 5)
	go receive(ch, &wg)
	wg.Wait()
}
