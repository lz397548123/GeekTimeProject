package once4

import "testing"

func Test_Once(t *testing.T) {
	var once Once
	once.doSlow()
}
