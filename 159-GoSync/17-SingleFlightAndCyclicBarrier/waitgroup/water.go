package main

import (
	"context"
	"golang.org/x/sync/semaphore"
	"sync"
)

// 定义水分子合成的辅助数据结构
type H2O struct {
	semaH *semaphore.Weighted // 氢原子的信号量
	semaO *semaphore.Weighted // 氧原子的信号量
	wg    sync.WaitGroup      //将循环栅栏替换成WaitGroup
}

func New() *H2O {
	var wg sync.WaitGroup
	wg.Add(3)

	return &H2O{
		semaH: semaphore.NewWeighted(2), // 氢原子需要两个
		semaO: semaphore.NewWeighted(1), // 氧原子需要一个
		wg:    wg,
	}
}

func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	h2o.semaH.Acquire(context.Background(), 1)

	releaseHydrogen() // 输出H

	// 标记自己已达到，等待其它goroutine到达
	h2o.wg.Done()
	h2o.wg.Wait()

	h2o.semaH.Release(1) // 释放氢原子空槽
}

func (h2o *H2O) oxygen(releaseOxygen func()) {
	h2o.semaO.Acquire(context.Background(), 1)

	releaseOxygen() // 输出O

	// 标记自己已达到，等待其它goroutine到达
	h2o.wg.Done()
	h2o.wg.Wait()
	// 都到达后重置
	h2o.wg.Add(3)

	h2o.semaO.Release(1) // 释放氧原子空槽
}
