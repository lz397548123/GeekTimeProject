package mapreduce

import "testing"

func asStream(done <-chan struct{}) <-chan any {
	s := make(chan any)
	values := []int{1, 2, 3, 4, 5}

	go func() {
		defer close(s)
		for _, v := range values { // 从数组生成
			select {
			case <-done:
				return
			case s <- v:
			}
		}
	}()

	return s
}

func TestMapReduce(t *testing.T) {
	in := asStream(nil)

	// map操作：乘以10
	mapFn := func(v any) any {
		return v.(int) * 10
	}

	// reduce操作：对map的结果进行累加
	reduceFn := func(r, v any) any {
		return r.(int) + v.(int)
	}

	sum := reduce(mapChan(in, mapFn), reduceFn) // 返回累加结果
	t.Log(sum)
}
