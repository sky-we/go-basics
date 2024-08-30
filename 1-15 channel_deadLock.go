package main

import "fmt"

//一、 无缓冲通道要求发送和接收操作必须同时进行，否则操作将会阻塞。以下是几种常见的死锁场景：

//1 没有接收者, 如果在向无缓冲通道发送数据时，没有 goroutine 准备接收数据，程序会因为发送操作无法完成而死锁。

func NoReceiver() {
	ch := make(chan int)

	ch <- 1 // 这里会死锁，因为没有其他 goroutine 接收数据
}

//2 主函数提前退出, 即使你启动了 goroutine 进行接收，如果主函数在接收完成前退出，可能会导致死锁或程序异常结束。

func MainExit() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	// 主函数提前退出，导致上面的 goroutine 没有机会运行
}

//3 互相等待, 如果两个 goroutine 互相等待对方发送或接收数据，也会导致死锁。

func WaitEachOther() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- <-ch2 // 等待从 ch2 接收数据后，再发送到 ch1
	}()

	go func() {
		ch2 <- <-ch1 // 等待从 ch1 接收数据后，再发送到 ch2
	}()

	// 两个 goroutine 互相等待，导致死锁
}

// 二、带缓冲通道允许一定数量的值存储在缓冲区中而不立即阻塞发送操作。但即使有缓冲，仍可能发生死锁。

//1 缓冲区满，没有接收者, 如果通道的缓冲区满了，但没有 goroutine 接收数据，程序会因为发送操作无法完成而死锁。

func BufferFullNoReceiver() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	ch <- 3 // 缓冲区满，没有接收者，导致死锁
}

//2 缓冲区空，没有发送者,当通道的缓冲区为空时，接收操作会阻塞。如果没有发送者，接收操作会导致死锁。

func BufferEmptyNoSender() {
	ch := make(chan int, 2)

	go func() {
		fmt.Println(<-ch)
	}()

	fmt.Println(<-ch) // 缓冲区为空，没有发送者，导致死锁
}

//3 忘记关闭通道, 如果发送方完成了所有的发送操作，但没有关闭通道，接收方可能会一直等待更多的数据，导致死锁。

func UnClosedChannel() {
	ch := make(chan int, 2)

	go func() {
		ch <- 1
		ch <- 2
		// ch.Close() // 如果忘记关闭通道，接收者会死锁
	}()

	for data := range ch {
		fmt.Println(data)
	}
}

/* 如何避免死锁?
合理的 goroutine 同步：确保每一个发送操作都有相应的接收操作，同时确保 goroutine 的生命周期管理得当，防止主函数提前退出。
关闭通道：发送方在完成所有发送操作后应关闭通道，通知接收方所有数据已经发送完毕。
使用 select 语句：在需要处理多路通道操作时，使用 select 可以避免因为某一个操作阻塞导致的死锁。
使用 sync.WaitGroup 等同步原语：对于复杂的 goroutine 协同操作，可以使用 sync.WaitGroup 等同步工具来确保所有 goroutine 的协同工作。
总结: 死锁通常发生在 goroutine 没有正确同步时，或者通道没有正确使用时。无缓冲通道要求发送和接收必须严格同步，而带缓冲通道虽然提供了缓冲区，但仍可能因为缓冲区满或空等原因导致死锁。在编写并发代码时，理解这些情况并小心设计代码逻辑是避免死锁的关键。
*/
