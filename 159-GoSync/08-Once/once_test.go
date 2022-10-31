package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"testing"
)

func Test_Once_Use(t *testing.T) {
	var once sync.Once

	// 第一个初始化函数
	f1 := func() {
		fmt.Println("in f1")
	}
	once.Do(f1) // 打印出 in f1

	// 第二个初始化函数
	f2 := func() {
		fmt.Println("in f2")
	}
	once.Do(f2) // 无输出
}

func Test_Once_Deadlock(t *testing.T) {
	var once sync.Once
	once.Do(func() {
		once.Do(func() {
			fmt.Println("初始化")
		})
	})
}

func Test_Once_NoInit(t *testing.T) {
	var once sync.Once
	var googleConn net.Conn // 到Google网站的一个连接

	once.Do(func() {
		// 建立到google.com的连接，有可能因为网络的原因，googleConn并没有建立成功，此时它的值为nil
		googleConn, _ = net.Dial("tcp", "google.com:80")
	})
	// 发送http请求
	googleConn.Write([]byte("GET / HTTP/1.1\r\nHost: google.com\r\n Accept: */*\r\n\r\n"))
	io.Copy(os.Stdout, googleConn)
}
