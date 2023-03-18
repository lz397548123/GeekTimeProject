package main

import "context"

type CyclicBarrier interface {
	// 等待所有的参与者到达，如果被ctx.Done()中断，会返回ErrBrokenBarrier
	Await(ctx context.Context) error
	// 重置循环栅栏到初始化状态。如果当前有等待者，那么它们会返回ErrBrokenBarrier
	Reset()
	// 返回当前等待者的数量
	GetNumberWaiting() int
	// 参与者的数量
	GetParties() int
	// 循环栅栏是否处于中断状态
	IsBroken() bool
}
