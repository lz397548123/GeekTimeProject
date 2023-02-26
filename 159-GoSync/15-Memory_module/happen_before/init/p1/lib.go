package p1

import (
	"GeekTimeProject/159-GoSync/15-Memory_module/happen_before/init/p2"
	"GeekTimeProject/159-GoSync/15-Memory_module/happen_before/init/trace"
	"fmt"
)

var V1_p1 = trace.Trace("init v1_p1", p2.V1_p2)
var V2_p1 = trace.Trace("init v2_p1", p2.V2_p2)

func init() {
	fmt.Println("init func in p1")
}
