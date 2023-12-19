package main

import (
	"fmt"
)

func main() {
	// Initialize a slice
	slice := []string{"apple", "banana", "cherry", "date", "elderberry"}

	// Iterating over the slice using a for loop
	fmt.Println("Iterating with for loop:")
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, slice[i])
	}

	// Iterating over the slice using range
	fmt.Println("\nIterating with range:")
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
}
