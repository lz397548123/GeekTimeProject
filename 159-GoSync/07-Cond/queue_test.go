package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

// Commands 用于产生出队，入队命令
func Commands(N int, random bool) []int {
	if N%2 != 0 {
		panic("will deadlock!")
	}
	// 0表示入队，1表示出队
	commands := make([]int, N)
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			commands[i] = 1
		}
	}

	if random {
		for i := len(commands) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			commands[i], commands[j] = commands[j], commands[i]
		}
	}
	return commands
}

func TestQueue(t *testing.T) {
	var wg sync.WaitGroup
	// 容量为5的阻塞队列
	q := NewQueue(5)

	// 生成随机命令
	for i, cmd := range Commands(20, true) {
		wg.Add(1)

		// 0表示入队，1表示出队
		if cmd == 0 {
			go func(id int) {
				defer wg.Done()
				q.Enqueue(id)
			}(i)
		} else {
			go func(id int) {
				defer wg.Done()
				q.Dequeue()
			}(i)
		}
	}

	wg.Wait()

	// 输出操作日志
	fmt.Println(q)
}
