package mutex2

import (
	"sync/atomic"
)

// 给新人机会、多给些机会 2011年6月30日

func runtime_Semacquire(s *uint32) {}
func runtime_Semrelease(s *uint32) {}

type Mutex struct {
	state int32
	sema  uint32
}

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWorken
	mutexWaiterShift = iota
)

func (m *Mutex) Lock() {
	// Fast path: 幸运case，能够直接获取到锁
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}

	awoke := false
	for {
		old := m.state
		new := old | mutexLocked // 新状态加锁
		if old&mutexLocked != 0 {
			new = old + 1<<mutexWaiterShift // 等待数量加1
		}
		if awoke {
			// goroutine是被唤醒的
			// 新状态清除唤醒标志 &^ 按位置零
			new &^= mutexWorken
		}
		if atomic.CompareAndSwapInt32(&m.state, old, new) { // 设置新状态
			if old&mutexLocked == 0 { // 锁原状态未加锁
				break
			}

			runtime_Semacquire(&m.sema) // 请求信号量
			awoke = true
		}
	}
}

func (m *Mutex) Unlock() {
	// Fast path: drop lock bit.
	new := atomic.AddInt32(&m.state, -mutexLocked) // 去掉锁标志
	if (new+mutexLocked)&mutexLocked == 0 {        // 本来就没有加锁
		panic("sync: unlock of unlocked mutex")
	}

	old := new
	for {
		if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWorken) != 0 {
			return
		}
		new = (old - 1<<mutexWaiterShift) | mutexWorken // 新状态，准备唤醒goroutine
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			runtime_Semrelease(&m.sema)
			return
		}
		old = m.state
	}
}
