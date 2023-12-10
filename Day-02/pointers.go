package main

import "fmt"

func main() {
	fmt.Println("map:")
	var s int = 10
	var t *int = &s // pointer to s

	fmt.Println("s=", s)   // prints the value of s
	fmt.Println("t=", t)   // prints the memory address of s
	fmt.Println("*t=", *t) // prints the value at the memory address stored in t (dereferencing)

}
