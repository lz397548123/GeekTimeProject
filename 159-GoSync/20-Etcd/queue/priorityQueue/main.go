package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	recipe "github.com/coreos/etcd/contrib/recipes"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	addr      = flag.String("addr", "http://42.193.109.34:2379", "etcd address")
	queueName = flag.String("name", "my-test-queue", "queue name")
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

	// 创建/获取队列
	q := recipe.NewPriorityQueue(cli, *queueName)

	// 从命令行读取命令
	consoleScanner := bufio.NewScanner(os.Stdin)
	for consoleScanner.Scan() {
		action := consoleScanner.Text()
		items := strings.Split(action, " ")
		switch items[0] {
		case "push": // 加入队列
			if len(items) != 3 {
				fmt.Println("must set value and priority to push")
				continue
			}
			pr, err := strconv.Atoi(items[2]) // 读取优先级
			if err != nil {
				fmt.Println("must set uint16 as priority", err)
				continue
			}
			q.Enqueue(items[1], uint16(pr)) // 入队
		case "pop": // 从队列弹出
			v, err := q.Dequeue() // 出队
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(v) // 输出出队的元素
		case "quit", "exit": // 退出
			return
		default:
			fmt.Println("unkonwn action")
		}
	}
}
