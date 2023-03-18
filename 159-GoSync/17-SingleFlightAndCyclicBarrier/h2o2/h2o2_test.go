package h2o2

import (
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestH2O2Factory(t *testing.T) {
	// 用来存放双氧水分子的channel
	var ch chan string
	releaseHydrogen := func() {
		ch <- "H"
	}
	releaseOxygen := func() {
		ch <- "O"
	}

	// 400个原子，400个goroutine，每个goroutine并发的产生一个原子
	var N = 100
	ch = make(chan string, N*4)

	h2o2 := New()

	// 用来等待所有的goroutine完成
	var wg sync.WaitGroup
	wg.Add(N * 4)

	// 200 个氢原子goroutine
	for i := 0; i < 2*N; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			h2o2.hydrogen(releaseHydrogen)
			wg.Done()
		}()
	}

	// 200个氧原子goroutine
	for i := 0; i < 2*N; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			h2o2.oxygen(releaseOxygen)
			wg.Done()
		}()
	}

	// 等待所有的goroutine执行完
	wg.Wait()

	// 结果中肯定是400个原子
	if len(ch) != N*4 {
		t.Fatalf("expect %d atom but got %d", N*4, len(ch))
	}

	// 每四个原子一组，分别进行检查。要求这一组原子中必须包含两个氢原子和两个氧原子，这样才能组成一个双氧水分子
	var s = make([]string, 4)
	for i := 0; i < N; i++ {
		s[0] = <-ch
		s[1] = <-ch
		s[2] = <-ch
		s[3] = <-ch
		sort.Strings(s)

		h2o2Str := s[0] + s[1] + s[2] + s[3]
		if h2o2Str != "HHOO" {
			t.Fatalf("expect a water molecule but got %s", h2o2Str)
		}
	}
}
