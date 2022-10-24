package main

import (
	"fmt"
	"sync"
)

// Go的issue 28123
type TestStruct struct {
	Wait sync.WaitGroup
}

func main() {
	w := sync.WaitGroup{}
	w.Add(1)
	t := &TestStruct{
		Wait: w,
	}

	t.Wait.Done()
	fmt.Println("Finished")
}
