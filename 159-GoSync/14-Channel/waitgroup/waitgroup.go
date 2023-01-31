package waitgroup

import (
	"fmt"
	"time"
)

func job(index int) {
	// 耗时任务
	time.Sleep(time.Second)
	fmt.Printf("任务：%d已完成\n", index)
}

func work(index int, limit chan struct{}) {
	limit <- struct{}{}
	job(index)
	<-limit
}
