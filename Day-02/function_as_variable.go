package main

import "fmt"

func main() {
	fmt.Println("function as variable:")
	var v func(int) int
	v = func(x int) int { return x * x }
	result := v(5) // result is 25

	fmt.Println("result=", result)
}
