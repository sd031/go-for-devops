package main

import "fmt"

func main() {
	fmt.Println("map:")
	var s int = 10
	var t *int = &s // pointer to s

	fmt.Println("s=", s)
	fmt.Println("t=", t)
}
