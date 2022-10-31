package muOnce

import (
	"fmt"
	"testing"
)

func Test_MuOnce(t *testing.T) {
	fmt.Println("Hello, playground")
	m := new(MuOnce)
	fmt.Println(m.strings())
	fmt.Println(m.strings())
}
