package main

import "fmt"

// Function to modify the value at a pointer
func modifyValue(ptr *int) {
	*ptr = 30
}

func main() {
	// Dereferencing a Pointer
	a := 20
	ptrA := &a
	fmt.Println("Original value of a:", a)
	fmt.Println("Dereferencing ptrA gives:", *ptrA)

	// Changing the Value at a Pointer
	*ptrA = 25
	fmt.Println("After changing, a is now:", a)

	// Comparing Pointers
	b := 20
	ptrB := &b
	fmt.Println("ptrA and ptrB point to the same address:", ptrA == ptrB)
	p1 := &a
	p2 := &b
	p3 := &a

	fmt.Println("p1 and p2 point to the same address?:", p1 == p2) // false
	fmt.Println("p1 and p2 to the same address?:", p1 == p3)       // true

	// Passing Pointers to Functions
	fmt.Println("Value of a before passing to function:", a)
	modifyValue(ptrA)
	fmt.Println("Value of a after passing to function:", a)
}
