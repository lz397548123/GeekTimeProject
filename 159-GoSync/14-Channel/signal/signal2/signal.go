package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func doCleanup(closed chan struct{}) {
	time.Sleep(time.Minute)
	close(closed)
}

func main() {
	closing := make(chan struct{})
	closed := make(chan struct{})

	go func() {
		// 模拟业务逻辑
		for {
			select {
			case <-closing:
				return
			default:
				// ...... 业务计算
				time.Sleep(100 * time.Minute)
			}
		}
	}()

	// 处理CTRL+C等中断信号
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	close(closing)
	// 执行退出之前的清理动作
	go doCleanup(closed)

	select {
	case <-closed:
	case <-time.After(time.Second):
		fmt.Println("清理超时，不等了")
	}

	fmt.Println("优雅退出")
}
