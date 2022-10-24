package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNegativeNumber(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)

	wg.Add(-10) // 将-10作为参数调用Add，计数值被设置为0

	wg.Add(-1) // 将-1作为参数调用Add，如果加上-1计数值就会变为负数。这是不对的，所以会触发
}

func TestOverDone(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	wg.Done()

	wg.Done()
}

func TestWaitGroup_Add1(t *testing.T) {
	dosomething := func(milliseconds time.Duration, wg *sync.WaitGroup) {
		duration := milliseconds * time.Millisecond
		time.Sleep(duration) // 故意sleep一段时间

		wg.Add(1)
		fmt.Println("后台执行，duration:", duration)
		wg.Done()
	}

	var wg sync.WaitGroup

	go dosomething(100, &wg) // 启动第一个goroutine
	go dosomething(110, &wg) // 启动第二个goroutine
	go dosomething(120, &wg) // 启动第三个goroutine
	go dosomething(130, &wg) // 启动第四个goroutine

	wg.Wait() // 主goroutine等待完成
	fmt.Println("Done")
}

func TestWaitGroup_Add2(t *testing.T) {
	dosomething := func(milliseconds time.Duration, wg *sync.WaitGroup) {
		duration := milliseconds * time.Millisecond
		time.Sleep(duration) // 故意sleep一段时间

		fmt.Println("后台执行，duration:", duration)
		wg.Done()
	}

	var wg sync.WaitGroup
	wg.Add(4) // 预先设定WaitGroup的计数值

	go dosomething(100, &wg) // 启动第一个goroutine
	go dosomething(110, &wg) // 启动第二个goroutine
	go dosomething(120, &wg) // 启动第三个goroutine
	go dosomething(130, &wg) // 启动第四个goroutine

	wg.Wait() // 主goroutine等待完成
	fmt.Println("Done")
}

func TestWaitGroup_Add3(t *testing.T) {
	dosomething := func(milliseconds time.Duration, wg *sync.WaitGroup) {
		wg.Add(1) // 计数值加1，在启动goroutine
		go func() {
			duration := milliseconds * time.Millisecond
			time.Sleep(duration) // 故意sleep一段时间

			fmt.Println("后台执行，duration:", duration)
			wg.Done()
		}()
	}
	var wg sync.WaitGroup

	dosomething(100, &wg) // 调用方法，把计数值加1，并启动任务goroutine
	dosomething(110, &wg) // 调用方法，把计数值加1，并启动任务goroutine
	dosomething(120, &wg) // 调用方法，把计数值加1，并启动任务goroutine
	dosomething(130, &wg) // 调用方法，把计数值加1，并启动任务goroutine

	wg.Wait() // 主goroutine等待，代码逻辑保证了四次Add(1)都已经执行完了
	fmt.Println("Done")
}

func TestWaitGroup_Reusing(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done() // 计数器减1
		wg.Add(1) // 计数值加1
	}()
	wg.Wait() // 主goroutine等待，有可能和第7行并发执行
}
