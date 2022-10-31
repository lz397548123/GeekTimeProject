package once3

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

// Once 是一个将扩展的sync.Once类型，提供了一个Done方法
type Once struct {
	sync.Once
}

// Done 返回此Once是否执行过
// 如果执行过则返回true
// 如果没有执行过或正在执行，返回false
func (o *Once) Done() bool {
	return atomic.LoadUint32((*uint32)(unsafe.Pointer(&o.Once))) == 1
}
