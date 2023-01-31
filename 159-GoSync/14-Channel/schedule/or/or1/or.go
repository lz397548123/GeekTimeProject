package or1

func or(channels ...<-chan any) <-chan any {
	// 特殊情况，只有零个或者1个chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan any)
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2: // 2个也是一种特殊情况
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default: // 超过两个，二分法递归处理
			m := len(channels) / 2
			select {
			case <-or(channels[:m]...):
			case <-or(channels[m:]...):
			}
		}
	}()

	return orDone
}
