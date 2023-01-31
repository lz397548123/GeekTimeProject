package or2

import (
	"testing"
	"time"
)

func sig(after time.Duration) <-chan any {
	c := make(chan any)
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func TestOr(t *testing.T) {
	start := time.Now()

	<-or(
		sig(10*time.Second),
		sig(20*time.Second),
		sig(30*time.Second),
		sig(40*time.Second),
		sig(50*time.Second),
		sig(01*time.Second),
	)

	t.Logf("done afer %v", time.Since(start))
}
