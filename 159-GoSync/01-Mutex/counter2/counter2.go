package main

import (
	"fmt"
	"sync"
)

/*
在这个例子中，我们创建了 10 个 goroutine，同时不断地对一个变量(count)进行加 1 操作，每个
goroutine 负责执行 10 万次的加 1 操作，我们期望的最后计数的结果是 10 * 100000 = 1000000 (一百万)。
*/

func main() {
	// 互斥锁保护计数器
	var mu sync.Mutex
	// 计数器的值
	var count = 0

	// 辅助变量，用来确认所有的goroutine都完成
	var wg sync.WaitGroup
	wg.Add(10)

	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 累加10万次
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	// 等待10哥goroutine完成
	wg.Wait()
	fmt.Println(count)
}
