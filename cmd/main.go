package main

import (
	. "async/cache"
	"fmt"
	"sync"
	"time"
)

const (
	k1   = "key1"
	step = 7
)

func main() {
	var wg sync.WaitGroup
	cache := NewCache()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Increase(k1, step)
			time.Sleep(time.Microsecond * 100)
		}()
	}
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Set(k1, step*i)
			time.Sleep(time.Millisecond * 100)
		}()
	}

	wg.Wait()
	fmt.Println(cache.Get(k1))
}

/*
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
*/
