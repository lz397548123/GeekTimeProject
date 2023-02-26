package main

import (
	"fmt"
)

var a string

func f() {
	fmt.Println(a)
}

func hello() {
	a = "hello world"
	go f()
}

func main() {
	hello()
	//time.Sleep(time.Second * 1) // 不睡眠有可能输出不出来
}
