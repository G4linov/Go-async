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
	done := make(chan struct{})

	for i := 0; i < 10; i++ {
		semaphore <- i
		go func() {
			defer func() {
				msg := <-semaphore
				fmt.Print(msg)
			}()
			cache.Increase(k1, step)
		}()
	}
	for i := 0; i < 10; i++ {
		semaphore <- i
		go func(i int) {
			defer func() {
				msg := <-semaphore
				fmt.Print(msg)
			}()
			cache.Set(k1, step*i)
		}(i)
	}
	for len(semaphore) > 0 {
		time.Sleep(time.Millisecond * 10)
	}
	go func() {
		time.Sleep(time.Millisecond * 1000)
		done <- struct{}{}
	}()
L:
	for {
		select {
		case <-done:
			fmt.Println("done")
			break L
		default:
			fmt.Println("waiting")
			time.Sleep(time.Millisecond * 100)
		}
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
