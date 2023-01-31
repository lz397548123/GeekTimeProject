package stream

func asStream(done <-chan struct{}, values ...any) <-chan any {
	s := make(chan any) // 创建一个unbuffered的channel
	go func() {         // 启动一个goroutine，往s中塞数据
		defer close(s)             // 退出时关闭chan
		for _, v := range values { // 遍历数组
			select {
			case <-done:
				return
			case s <- v: // 将数据元素塞入到chan中
			}
		}
	}()
	return s
}

func takeN(done <-chan struct{}, valueStream <-chan any, num int) <-chan any {
	takeStream := make(chan any) // 创建输出流
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ { // 只读取前num个元素
			select {
			case <-done:
				return
			case takeStream <- <-valueStream: // 从输入流中读取数据
			}
		}
	}()
	return takeStream
}

func takeFn(done <-chan struct{}, valueStream <-chan any, fn func(item any) bool) <-chan any {
	takeStream := make(chan any) // 创建输出流
	go func() {
		defer close(takeStream)
		for v := range valueStream {
			if fn(v) {
				select {
				case <-done:
					return
				case takeStream <- v:
				}
			}
		}
	}()
	return takeStream
}

func takeWhile(done <-chan struct{}, valueStream <-chan any, fn func(item any) bool) <-chan any {
	takeStream := make(chan any) // 创建输出流
	go func() {
		defer close(takeStream)
		for v := range valueStream {
			if !fn(v) {
				return
			}

			select {
			case <-done:
				return
			case takeStream <- v:
			}
		}
	}()
	return takeStream
}

func skipN(done <-chan struct{}, valueStream <-chan any, num int) <-chan any {
	takeStream := make(chan any) // 创建输出流
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			<-valueStream
		}
		for v := range valueStream {
			select {
			case <-done:
				return
			case takeStream <- v:
			}
		}
	}()
	return takeStream
}

func skipFn(done <-chan struct{}, valueStream <-chan any, fn func(item any) bool) <-chan any {
	takeStream := make(chan any) // 创建输出流
	go func() {
		defer close(takeStream)
		for v := range valueStream {
			if fn(v) {
				continue
			}
			select {
			case <-done:
				return
			case takeStream <- v:
			}
		}
	}()
	return takeStream
}

func skipWhile(done <-chan struct{}, valueStream <-chan any, fn func(item any) bool) <-chan any {
	takeStream := make(chan any) // 创建输出流
	go func() {
		defer close(takeStream)

		for { // 找到第一个不满足的
			v, ok := <-valueStream
			if !ok { // 没数据直接返回
				return
			}

			if !fn(v) {
				takeStream <- v
				break
			}
		}

		for v := range valueStream {
			select {
			case <-done:
				return
			case takeStream <- v:
			}
		}
	}()
	return takeStream
}
