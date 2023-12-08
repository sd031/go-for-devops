package main

import "fmt"

func main() {
	x := 10

	fmt.Println("Basic if:")
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	fmt.Println("if..else:")

	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}

	fmt.Println("if-else if-else:")
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x == 10 {
		fmt.Println("x is exactly 10")
	} else {
		fmt.Println("x is less than 10")
	}

	fmt.Println("With Initialization Statement:")
	if y := 20; y > 10 {
		fmt.Println("y is greater than 10")
	}
	// Note: y is not accessible outside the if block

}
