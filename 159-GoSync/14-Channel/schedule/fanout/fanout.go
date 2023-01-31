package fanout

func fanOut(ch <-chan any, out []chan any, async bool) {
	go func() {
		defer func() { // 退出时关闭所有的输出chan
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		for v := range ch { // 从输入chan中读取数据
			value := v
			for i := 0; i < len(out); i++ {
				index := i
				if async { // 异步
					go func() {
						out[index] <- value // 放入到输出chan中，异步方式
					}()
				} else {
					out[index] <- value // 放入到输出chan中，同步方式
				}
			}
		}
	}()
}
