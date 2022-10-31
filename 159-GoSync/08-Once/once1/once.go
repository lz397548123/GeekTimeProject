package once1

import "sync/atomic"

/*
这确实是一种实现方式，但是，这个实现有一个很大的问题，就是如果参数 f 执行很慢的话，
后续调用 Do 方法的 goroutine 虽然看到 done 已经设置为执行过了，但是获取某些初
始化资源的时候可能会得到空的资源，因为 f 还没有执行完。
*/

type Once struct {
	done uint32
}

func (o *Once) Do(f func()) {
	if !atomic.CompareAndSwapUint32(&o.done, 0, 1) {
		return
	}
	f()
}
