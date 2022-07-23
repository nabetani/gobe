package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

type proc = func(a int32) int32

func test(num int32) int32 {
	m := make([]proc, 0, num)
	for i := int32(0); i < num; i++ {
		m = append(m, func(a int32) int32 {
			x := a*7 + 11
			y := x*13 + 15
			z := y*17 + 19
			w := z*23 + 29
			t := x ^ (x << 11)
			return ((w ^ (w >> 19)) ^ (t ^ (t >> 8))) % num
		})
	}
	sum := int32(0)

	for i := int32(0); i < num; i++ {
		b := make([]bool, num)
		sum += func(s int32) int32 {
			for {
				if b[s] {
					return s
				}
				b[s] = true
				s = m[s](i)
			}
		}(i)
		sum %= 1 << 24
	}
	return sum
}

func main() {
	num, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		panic(err)
	}
	t0 := time.Now()
	r := test(int32(num))
	tick := time.Now().Sub(t0)
	ms := float64(tick.Microseconds()) * 1e-3
	fmt.Printf("r:%d, Go version:%q, tick:%.2fms\n", r, runtime.Version(), ms)
}
