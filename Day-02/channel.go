package main

import "fmt"

func main() {
	fmt.Println("channel:")
	u := make(chan int) // unbuffered channel of integers
	fmt.Println("u=", u)
}
