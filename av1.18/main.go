package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

type proc = func(a uint32) uint32

func test(num int) int {
	chans := make([]chan int, num+1)
	for i := 0; i < len(chans); i++ {
		chans[i] = make(chan int)
	}
	count := num
	for i := 0; i < num; i++ {
		go func(ch []chan int, ix int) {
			for c := 0; c < count; c++ {
				v := <-ch[0]
				ch[1] <- (v + 1 + ix + c) % num
			}
		}(chans[i:i+2], i)
	}
	go func() {
		for c := 0; c < count; c++ {
			chans[0] <- (c + 1)
		}
	}()
	sum := 0
	for c := 0; c < count; c++ {
		sum += <-chans[num]
	}
	return sum
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
	fmt.Printf("r:%d, Go version:%q, tick:%.2fms\n", r, runtime.Version(), ms)
}
