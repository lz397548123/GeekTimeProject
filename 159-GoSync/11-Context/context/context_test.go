package context

import (
	"context"
	"testing"
	"time"
)

func TestWithValue(t *testing.T) {
	ctx := context.Background()
	ctx = context.TODO()
	ctx = context.WithValue(ctx, "key1", "0001")
	ctx = context.WithValue(ctx, "key2", "0002")
	ctx = context.WithValue(ctx, "key3", "0003")
	ctx = context.WithValue(ctx, "key4", "0004")
	t.Log(ctx.Value("key1"))
}

func TestCheckContextDoneClose(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer func() {
			t.Log("goroutine exit")
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}

func TestThinkingQuestion(t *testing.T) {
	watchParent := func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // 取出值即说明是结束信号
				t.Log("收到信号，父context的协程推出，time=", time.Now().Unix())
				return
			default:
				t.Log("父context的协程监控中,time=", time.Now().Unix())
				time.Sleep(1 * time.Second)
			}
		}
	}

	watchChild := func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // 取出值即说明是结束信号
				t.Log("收到信号，子context的协程推出，time=", time.Now().Unix())
				return
			default:
				t.Log("子context的协程监控中,time=", time.Now().Unix())
				time.Sleep(1 * time.Second)
			}
		}
	}

	// 父context(利用根context得到)
	ctx, cancel := context.WithCancel(context.Background())

	// 父context的子协程
	go watchParent(ctx)

	// 子context，注意：这里虽然也返回了cancel的函数对象，但是未使用
	valueCtx, _ := context.WithCancel(ctx)

	go watchChild(valueCtx)

	t.Log("现在开始等待3秒，time=", time.Now().Unix())
	time.Sleep(3 * time.Second)

	// 调用cancel()
	t.Log("等待3秒结束， 调用cancel()函数")
	cancel()

	// 再等待5秒看输出，可以发现父context的子协程和子context的子协程都会被结束掉
	time.Sleep(5 * time.Second)
	t.Log("最终结束,time=", time.Now().Unix())
}
