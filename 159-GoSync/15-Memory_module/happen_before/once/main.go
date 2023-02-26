package main

import "sync"

var s string
var once sync.Once

func foo() {
	s = "hello world"
}

func main() {
	once.Do(foo)
	print(s)
}
