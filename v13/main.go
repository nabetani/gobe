package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	t0 := time.Now()
	tick := time.Now().Sub(t0)
	ms := float64(tick/time.Microsecond) * 1e-3
	fmt.Printf("Go version:%q, tick:%.2fms\n", runtime.Version(), ms)
}
