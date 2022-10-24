package main

import (
	"fmt"
	"sync"
)

func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("int bar")
	l.Unlock()
}

func Foo(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}

func main() {
	l := &sync.Mutex{}
	Foo(l)
}
