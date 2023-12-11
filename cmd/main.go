package main

import (
	"fmt"
	"time"
)

func printNumbers(m, n int) {
	for i := m; i < n; i++ {
		fmt.Println(i)
	}
}

func main() {
	go printNumbers(0, 5)
	go fmt.Println("end of print")
	time.Sleep(time.Second)
}
