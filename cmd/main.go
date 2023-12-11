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
	printNumbers(0, 5)
	fmt.Println("end of print")
	time.Sleep(time.Second)
}
