package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func test(num int) int {
	var pre func(d int) []interface{}
	var get func(v interface{}) int
	pre = func(d int) []interface{} {
		if d <= 1 {
			return []interface{}{1}
		}
		return []interface{}{pre(d - 1), pre(d - 2)}
	}
	get = func(v interface{}) int {
		switch val := v.(type) {
		case int:
			return val
		case []interface{}:
			sum := 0
			for _, e := range val {
				sum += get(e)
			}
			return sum
		default:
			panic("logic error")
		}
	}
	v := pre(num)
	return get(v)
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
