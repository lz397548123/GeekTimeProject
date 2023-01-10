package channel

import (
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	var ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case v := <-ch:
			t.Log(v)
		}
	}
}

func TestRange(t *testing.T) {
	var ch = make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			t.Logf("传入数字:%d 时间:%d", i, time.Now().Unix())
			ch <- i
		}
	}()

	go func() {
		for v := range ch {
			t.Logf("传出数字:%d 时间:%d", v, time.Now().Unix())
		}
	}()

	time.Sleep(time.Second)
}

func TestRange2(t *testing.T) {
	var ch = make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for range ch {
	}

	v, ok := <-ch
	t.Log(v, ok)
}

func TestGoroutineLeak(t *testing.T) {
	process(time.Second)
}
