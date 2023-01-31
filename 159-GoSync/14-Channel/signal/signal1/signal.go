package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func doCleanup() {
	fmt.Println("连接关闭、文件close、缓存落盘")
}

func main() {
	go func() {
		fmt.Println("执行业务逻辑")
	}()

	// 处理CTRL+C等中断信号
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	// 执行退出之前的清理动作
	doCleanup()

	fmt.Println("优雅退出")
}
