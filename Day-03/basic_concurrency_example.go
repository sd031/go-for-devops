package main

import (
	"fmt"
	"time"
)

func printCounts(label string, count int, ch chan int) {
	for i := 0; i < count; i++ {
		ch <- i
		fmt.Println(label, i)
		time.Sleep(time.Millisecond * 500) // simulate work
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go printCounts("goroutine", 5, ch)

	for value := range ch {
		fmt.Println("Main received:", value)
	}
}
