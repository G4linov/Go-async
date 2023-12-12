package main

import (
	. "async/cache"
	"fmt"
	"time"
)

const (
	k1   = "key1"
	step = 7
)

func main() {
	semaphore := make(chan int, 4)
	cache := NewCache()

	for i := 0; i < 10; i++ {
		semaphore <- i
		go func() {
			defer func() {
				msg := <-semaphore
				fmt.Println(msg)
			}()
			cache.Increase(k1, step)
			time.Sleep(time.Microsecond * 100)
		}()
	}
	for i := 0; i < 10; i++ {
		semaphore <- i
		go func(i int) {
			defer func() {
				msg := <-semaphore
				fmt.Println(msg)
			}()
			cache.Set(k1, step*i)
			time.Sleep(time.Microsecond * 100)
		}(i)
	}
	for len(semaphore) > 0 {
		time.Sleep(time.Millisecond * 1000)
	}
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
