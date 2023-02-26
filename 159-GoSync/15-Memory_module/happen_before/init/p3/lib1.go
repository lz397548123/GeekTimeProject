package p3

import (
	"GeekTimeProject/159-GoSync/15-Memory_module/happen_before/init/trace"
	"fmt"
)

var V1_p3 = trace.Trace("init v1_p3", 3)
var V2_p3 = trace.Trace("init v2_p3", 3)

func init() {
	fmt.Println("init func in p3")
	V1_p3 = 300
	V2_p3 = 300
}
