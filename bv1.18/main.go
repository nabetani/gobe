package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"time"
)

func test(num int) float64 {
	a := make([]float64, num)
	for end := num; 0 < end; end-- {
		for i := end - 1; 3 <= i; i-- {
			a[i/2] = math.Mod((1+a[i]+math.Sin(a[i-1])+math.Cos(a[i-2]))/(1+a[i-3]*a[i-3]), 100)
		}
	}
	return a[0]
}

func main() {
	num, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		panic(err)
	}
	t0 := time.Now()
	r := test(int(num))
	tick := time.Now().Sub(t0)
	ms := float64(tick.Nanoseconds()) * 1e-6
	fmt.Printf("r:%f, n=%d, Go version:%q, tick:%.2fms\n", r, num, runtime.Version(), ms)
}
