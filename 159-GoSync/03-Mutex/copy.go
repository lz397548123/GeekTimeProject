package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	Counter int
}

// 这里Counter的参数是通过复制的方式传入的
func foo(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}

func main() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Counter++
	foo(c) // 复制锁
}
