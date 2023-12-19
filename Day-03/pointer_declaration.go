package main

import "fmt"

func main() {
	// Using var
	var ptr1 *int
	fmt.Println("1. Using var to declare a pointer:", ptr1)

	// Pointer to an Existing Variable
	var value int = 42
	ptr2 := &value
	fmt.Println("2. Pointer to an Existing Variable:", *ptr2)

	// Creating a Pointer with new
	ptr3 := new(int)
	*ptr3 = 100
	fmt.Println("3. Creating a Pointer with new:", *ptr3)

	// Pointer to Pointer
	ptr4 := &ptr3
	fmt.Println("4. Pointer to Pointer:", **ptr4)
}
