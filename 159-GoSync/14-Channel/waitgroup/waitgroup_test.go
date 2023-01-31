package waitgroup

import (
	"sync"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	limit := make(chan struct{}, 10)
	jobCount := 100
	for i := 1; i <= jobCount; i++ {
		go work(i, limit)
	}

	time.Sleep(20 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	jobCount := 100
	limit := 10
	for i := 1; i <= jobCount; i += limit {
		for j := i; j < i+limit && j <= jobCount; j++ {
			wg.Add(1)
			go func(item int) {
				defer wg.Done()
				job(item)
			}(j)
		}
		wg.Wait()
	}
}
