package main

import (
	. "async/cache"
	"fmt"
	"time"
)

func main() {
	c0 := make(chan int)
	go spinner(100 * time.Millisecond)
	go func() {
		var n3, n1, n2 = 0, 0, 1
		for i := 1; i <= 44; i++ {
			n3 = n1 + n2
			n1 = n2
			n2 = n3
		}
		c0 <- n3
	}()
	done := <-c0
	n := (45)
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", 44, done)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

	sad := make(map[string]int)

	cac := NewCache(sad)

	cac.Set("1st", 1)
	go cac.Increase("1st", 3)
	cac.Increase("1st", 2)

	fmt.Println(cac.Get("1st"))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
