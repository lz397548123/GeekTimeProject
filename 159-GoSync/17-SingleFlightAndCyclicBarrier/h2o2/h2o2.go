package h2o2

import (
	"context"
	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/semaphore"
)

// 定义双氧水分子合成的辅助数据结构
type H2O2 struct {
	semaH *semaphore.Weighted         // 氢原子的信号量
	semaO *semaphore.Weighted         // 氧原子的信号量
	b     cyclicbarrier.CyclicBarrier // 循环栅栏，用来控制合成
}

func New() *H2O2 {
	return &H2O2{
		semaH: semaphore.NewWeighted(2), // 氢原子需要两个
		semaO: semaphore.NewWeighted(2), // 氧原子需要两个
		b:     cyclicbarrier.New(4),     // 需要四个原子才能合成
	}
}

func (h2o2 *H2O2) hydrogen(releaseHydrogen func()) {
	h2o2.semaH.Acquire(context.Background(), 1)

	releaseHydrogen()                  // 输出H
	h2o2.b.Await(context.Background()) // 等待栅栏放行
	h2o2.semaH.Release(1)              // 释放氢原子空槽
}

func (h2o2 *H2O2) oxygen(releaseOxygen func()) {
	h2o2.semaO.Acquire(context.Background(), 1)

	releaseOxygen()                    // 输出O
	h2o2.b.Await(context.Background()) // 等待栅栏放行
	h2o2.semaO.Release(1)              // 释放氧原子空槽
}
