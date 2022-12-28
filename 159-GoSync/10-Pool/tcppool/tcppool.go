package tcppool

import (
	"fmt"
	"gopkg.in/fatih/pool.v2"
	"net"
)

func TcpPool() {
	// 工厂模式，提供创建连接的工厂方法
	factory := func() (net.Conn, error) {
		return net.Dial("tcp", "127.0.0.1:4000")
	}

	// 创建一个tcp池，提供厨师容量和最大容量以及工厂方法
	p, err := pool.NewChannelPool(5, 30, factory)
	if err != nil {
		_ = fmt.Errorf("%s", err)
	}

	// 获取一个连接
	conn, err := p.Get()
	if err != nil {
		_ = fmt.Errorf("%s", err)
	}

	// Close并不会真正关闭这个连接，而是把它放回池子，所以你不必显式地Put这个对象到池子中
	err = conn.Close()
	if err != nil {
		_ = fmt.Errorf("%s", err)
	}

	// 通过调用MarkUnusable，Close的时候就会真正关闭底层的tcp连接了
	if pc, ok := conn.(*pool.PoolConn); ok {
		pc.MarkUnusable()
		err = pc.Close()
		if err != nil {
			_ = fmt.Errorf("%s", err)
		}
	}
}
