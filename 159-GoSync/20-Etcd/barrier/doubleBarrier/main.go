package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	recipe "github.com/coreos/etcd/contrib/recipes"
	"log"
	"os"
	"strings"
)

var (
	addr        = flag.String("addr", "http://42.193.109.34:2379", "etcd address")
	barrierName = flag.String("name", "my-test-doubleBarrier", "barrier name")
	count       = flag.Int("c", 2, "")
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

	// 创建session
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()

	// 创建/获取栅栏
	b := recipe.NewDoubleBarrier(s1, *barrierName, *count)

	// 从命令行读取命令
	consoleScanner := bufio.NewScanner(os.Stdin)
	for consoleScanner.Scan() {
		action := consoleScanner.Text()
		items := strings.Split(action, " ")
		switch items[0] {
		case "enter": // 持有这个barrier
			b.Enter()
			fmt.Println("enter")
		case "leave": // 释放这个barrier
			b.Leave()
			fmt.Println("leave")
		case "quit", "exit": // 退出
			return
		default:
			fmt.Println("unkonwn action")
		}
	}
}
