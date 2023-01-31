package fanin

import "reflect"

func fanInReflect(chans ...<-chan any) <-chan any {
	out := make(chan any)

	go func() {
		defer close(out)
		// 构造SelectCases slice
		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		// 循环，从cases中选择一个可用的
		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok { // 此channel已经close
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface()
		}
	}()

	return out
}

func mergeTwo(a, b <-chan any) <-chan any {
	c := make(chan any)

	go func() {
		defer close(c)
		for a != nil || b != nil { // 只要还有可读的chan
			select {
			case v, ok := <-a:
				if !ok { // a关闭，设置为nil
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok { // b关闭，设置为nil
					b = nil
					continue
				}
				c <- v
			}
		}
	}()

	return c
}

func fanInRec(chans ...<-chan any) <-chan any {
	switch len(chans) {
	case 0:
		c := make(chan any)
		close(c)
		return c
	case 1:
		return chans[0]
	case 2:
		return mergeTwo(chans[0], chans[1])
	default:
		m := len(chans) / 2
		return mergeTwo(fanInRec(chans[:m]...), fanInRec(chans[m:]...))
	}
}
