package once3

import (
	"fmt"
	"testing"
	"time"
)

func TestOnce_Done(t *testing.T) {
	var flag Once
	fmt.Println(flag.Done()) // false

	flag.Do(func() {
		time.Sleep(time.Second)
	})

	fmt.Println(flag.Done()) // true
}
