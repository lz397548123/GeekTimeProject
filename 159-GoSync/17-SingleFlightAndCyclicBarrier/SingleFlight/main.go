package main

import "sync"

// 代表一个正在处理的请求，或者已经处理完的请求
type call struct {
	wg sync.WaitGroup

	// 这个字段代表处理完的值，在waitgroup完成之前只会写一次
	// waitgroup完成之后就读取这个值
	val any
	err error

	// 指示当call在处理时是否要忘记这个key
	forgotten bool
	dups      int
	chans     []chan<- any
}

// Group 代表一个singleflight对象
type Group struct {
	mu sync.Mutex       // protects m
	m  map[string]*call // lazily initialied
}

type Result struct {
	val    any
	err    error
	shared bool
}

func (g *Group) Do(key string, fn func() (any, error)) (v any, err error, shared bool) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok { // 如果已经存在相同的key
		c.dups++
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err, true
	}

	c := new(call) // 第一个请求，创建一个call
	c.wg.Add(1)
	g.m[key] = c // 加入到key mao中
	g.mu.Unlock()

	g.doCall(c, key, fn) // 调用方法
	return c.val, c.err, c.dups > 0
}

func (g *Group) doCall(c *call, key string, fn func() (any, error)) {
	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()

	if !c.forgotten { // 已调用完，删除这个key
		delete(g.m, key)
	}

	for _, ch := range c.chans {
		ch <- Result{c.val, c.err, c.dups > 0}
	}

	g.mu.Unlock()
}

func main() {

}
