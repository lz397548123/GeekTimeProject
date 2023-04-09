package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	recipe "github.com/coreos/etcd/contrib/recipes"
	"log"
	"os"
	"strings"
)

var (
	addr        = flag.String("addr", "http://42.193.109.34:2379", "etcd address")
	barrierName = flag.String("name", "my-test-barrier", "barrier name")
)

func main() {
	flag.Parse()

	// 解析etcd地址
	endpoints := strings.Split(*addr, ",")

	// 创建etcd的client
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 创建/获取栅栏
	b := recipe.NewBarrier(cli, *barrierName)

	// 从命令行读取命令
	consoleScanner := bufio.NewScanner(os.Stdin)
	for consoleScanner.Scan() {
		action := consoleScanner.Text()
		items := strings.Split(action, " ")
		switch items[0] {
		case "hold": // 持有这个barrier
			b.Hold()
			fmt.Println("hold")
		case "release": // 释放这个barrier
			b.Release()
			fmt.Println("released")
		case "wait": // 等待barrier被释放
			b.Wait()
			fmt.Println("after wait")
		case "quit", "exit": // 退出
			return
		default:
			fmt.Println("unkonwn action")
		}
	}
}
