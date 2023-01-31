package mapreduce

func mapChan(in <-chan any, fn func(any) any) <-chan any {
	out := make(chan any) // 创建一个输出chan
	if in == nil {        // 异常检查
		close(out)
		return out
	}

	go func() {
		defer close(out)
		for v := range in { // 从输入chan读取数据，执行业务操作，也就是map操作
			out <- fn(v)
		}
	}()

	return out
}

func reduce(in <-chan any, fn func(r, v any) any) any {
	if in == nil { // 异常检查
		return nil
	}

	out := <-in         // 先读取一个元素
	for v := range in { // 实现reduce的主要逻辑
		out = fn(out, v)
	}

	return out
}
