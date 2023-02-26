package p2

import (
	"GeekTimeProject/159-GoSync/15-Memory_module/happen_before/init/p3"
	"GeekTimeProject/159-GoSync/15-Memory_module/happen_before/init/trace"
	"fmt"
)

var V1_p2 = trace.Trace("init v1_p2", 2)
var V2_p2 = trace.Trace("init v2_p2", p3.V2_p3)

func init() {
	fmt.Println("init func in p2")
	V1_p2 = 200
}
