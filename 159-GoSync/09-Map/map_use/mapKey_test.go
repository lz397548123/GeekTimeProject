package map_use

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestMapKeyStruct(t *testing.T) {
	var m = make(map[mapKey]string)
	var key = mapKey{10}

	m[key] = "hello"
	fmt.Printf("m[key]=%s\n", m[key])

	// 修改key的字段的值后再次查询map，无法获取刚才add进去的值
	key.key = 100
	fmt.Printf("再次查询m[key]=%s\n", m[key])
}

func TestReturnTwoValue(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 0
	fmt.Printf("a=%d; b=%d\n", m["a"], m["b"])

	av, aExisted := m["a"]
	bv, bExisted := m["b"]
	fmt.Printf("a=%d, existed: %t; b=%d, existed: %t\n", av, aExisted, bv, bExisted)
}

func TestMapNoInit(t *testing.T) {
	var m map[int]int
	m[100] = 100
}

func TestMapNoInitGet(t *testing.T) {
	var m map[int]int
	fmt.Println(m[100])
}

type Counter struct {
	Website      string
	Start        time.Time
	PageCounters map[string]int
}

func TestCounter(t *testing.T) {
	var c Counter
	c.Website = "baidu.com"

	c.PageCounters["/"]++
}

func TestMapConcurrentPanic(t *testing.T) {
	var m = make(map[int]int, 10) // 初始化一个map
	go func() {
		for {
			m[1] = 1 // 设置key
		}
	}()

	go func() {
		for {
			_ = m[2] // 访问这个map
		}
	}()

	select {}
}

func Test111(t *testing.T) {
	s := "12341"
	t.Log(strings.Trim(s, "1"))
}
