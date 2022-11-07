package mutex3

import (
	"sync/atomic"
)

func runtime_Semacquire(s *uint32) {}
func runtime_canSpin(i int) bool   { return true }
func runtime_doSpin()              {}

type Mutex struct {
	state int32
	sema  uint32
}

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWorken
	mutexWaiterShift = iota
)

// 多给些机会 2015年2月---------
func (m *Mutex) Lock() {
	// Fast path: 幸运之路，正好获取到锁
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}

	awoke := false
	iter := 0
	for { // 不管是新来的请求锁的 goroutine，还是被唤醒的 goroutine，都不断尝试请求锁
		old := m.state            // 先保存当前锁的状态
		new := old | mutexLocked  // 新状态设置加锁标志
		if old&mutexLocked != 0 { // 锁还没被释放
			if runtime_canSpin(iter) { // 还可以自旋
				if !awoke && old&mutexWorken == 0 && old>>mutexWaiterShift != 0 &&
					atomic.CompareAndSwapInt32(&m.state, old, old|mutexWorken) {
					awoke = true
				}
				runtime_doSpin()
				iter++
				continue // 自旋，再次尝试请求锁
			}
			new = old + 1<<mutexWaiterShift
		}
		if awoke { // 唤醒状态
			if new&mutexWorken == 0 {
				panic("sync: inconsistent mutex state")
			}
			new &^= mutexWorken // 新状态清除唤醒标记
		}
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			if old&mutexLocked == 0 { // 旧状态锁已释放，新状态成功持有了锁，直接返回
				break
			}
			runtime_Semacquire(&m.sema)
			awoke = true // 被唤醒
			iter = 0
		}

	}
}
