package once4

import "sync"

type Once struct {
	m sync.Mutex
}

func (o *Once) doSlow() {
	o.m.Lock()
	defer o.m.Unlock()

	// 这里更新o指针的值！！！！！！，会导致上一行Unlock出错
	*o = Once{}
}
