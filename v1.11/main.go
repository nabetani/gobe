package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

type proc = func(a uint32) uint32

func test(num uint32) uint32 {
	m := make([]proc, 0, num)
	for ix := uint32(0); ix < num; ix++ {
		i := ix
		m = append(m, func(a uint32) uint32 {
			x := (i+a)*7 + 11
			y := (x+a)*13 + 15
			z := (y+i)*17 + 19
			w := (z+i+a)*23 + 29
			t := x ^ (x << 11)
			return ((w ^ (w >> 19)) ^ (t ^ (t >> 8))) % num
		})
	}
	sum := uint32(0)
	for i := uint32(0); i < num; i++ {
		sum += func(s0 uint32) uint32 {
			s := s0
			b := make([]bool, num)
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
	r := test(uint32(num))
	tick := time.Now().Sub(t0)
	ms := float64(tick.Nanoseconds()) * 1e-6
	fmt.Printf("r:%d, Go version:%q, tick:%.2fms\n", r, runtime.Version(), ms)
}
