package main

import (
	"fmt"
	"sync"
)

// 嵌入字段的方式

func main() {
	type Counter struct {
		mu    sync.Mutex
		Count uint64
	}
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.mu.Lock()
				counter.Count++
				counter.mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count)
}
