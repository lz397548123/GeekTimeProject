package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	recipe "github.com/coreos/etcd/contrib/recipes"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	addr     = flag.String("addr", "http://42.193.109.34:2379", "etcd address")
	lockName = flag.String("name", "my-test-lock", "lock name")
	action   = flag.String("rw", "w", "r means acquiring read lock, w means acquiring write lock")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	// 解析etcd地址
	endpoints := strings.Split(*addr, ",")

	// 创建一个etcd client
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
	m1 := recipe.NewRWMutex(s1, *lockName)

	// 从命令行读取命令
	consoleScanner := bufio.NewScanner(os.Stdin)
	for consoleScanner.Scan() {
		action := consoleScanner.Text()
		switch action {
		case "w": // 请求写锁
			testWriteLocker(m1)
		case "r": // 请求读锁
			testReadLocker(m1)
		default:
			fmt.Println("unkonwn action")
		}
	}
}

func testWriteLocker(m1 *recipe.RWMutex) {
	// 请求写锁
	log.Println("acquiring write lock")
	if err := m1.Lock(); err != nil {
		log.Fatal(err)
	}
	log.Println("acquired write lock")

	// 等待一段时间
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	// 释放写锁
	if err := m1.Unlock(); err != nil {
		log.Fatal(err)
	}
	log.Println("released write lock")
}

func testReadLocker(m1 *recipe.RWMutex) {
	// 请求读锁
	log.Println("acquiring read lock")
	if err := m1.RLock(); err != nil {
		log.Fatal(err)
	}
	log.Println("acquired read lock")

	// 等待一段时间
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	// 释放读锁
	if err := m1.RUnlock(); err != nil {
		log.Fatal(err)
	}
	log.Println("released read lock")
}
