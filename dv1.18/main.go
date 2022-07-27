package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func test(num int) int {
	var f func(int) int
	f = func(n int) int {
		if n < 2 {
			return n
		}
		ch := make(chan int)
		go func() { ch <- f(n - 1) }()
		go func() { ch <- f(n - 2) }()
		return <-ch + <-ch
	}
	return f(int(num))
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
	fmt.Printf("r:%d, n=%d, Go version:%q, tick:%.2fms\n", r, num, runtime.Version(), ms)
}
