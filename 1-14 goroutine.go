package main

import (
	"fmt"
	"sync"
	"time"
)

/*
并发：同时运行多个任务，但不一定在同一时刻运行
并行：同时运行多个任务，并且在同一时刻运行，需要多核CPU实现
为什么需要并发：
IO操作：通过并发可以充分利用CPU的空闲时间
死锁：使用同步避免共享资源死锁
扩展性：将大任务分解为小任务，分配给多线程执行，充分利用CPU
*/

var students = map[int64]*student{
	1: {
		id:      0,
		name:    "lw",
		age:     18,
		address: "公租房",
	},
	2: {
		id:      1,
		name:    "wj",
		age:     20,
		address: "湾湾超市",
	},
}

type student struct {
	id      int64
	name    string
	age     int
	address string
}

func getStudent(id int64) (*student, error) {
	time.Sleep(time.Millisecond * 100)
	return students[id], nil
}

func main() {
	start := time.Now()
	ids := []int64{1, 2}
	idAgeMap := make(map[int64]int)
	var (
		eg    sync.WaitGroup
		mutex sync.Mutex
	)
	for _, id := range ids {
		eg.Add(1)
		go func(id int64) {

			defer func() {
				// defer释放资源
				eg.Done()

			}()
			s, err := getStudent(id)
			if err != nil {
				panic("error")
			}
			mutex.Lock() // 互斥锁限制并发协程对共享资源的访问，避免了数据不一致和竞态条件
			idAgeMap[s.id] = s.age
			mutex.Unlock()
			fmt.Printf("student id is: %d, s is: [%+v]\n", s.id, s)
		}(id)

	}
	eg.Wait()
	end := time.Since(start)
	fmt.Println(end)
	fmt.Println(idAgeMap)
}
