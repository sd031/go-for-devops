package main

import (
	"fmt"
)

// modifySlice function modifies the slice it receives
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 999 // Modify the first element
	}
	fmt.Println("Inside modifySlice, after modification:", s)
}

func main() {
	// Create a slice
	originalSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice before modification:", originalSlice)

	// Pass the slice to the modifySlice function
	modifySlice(originalSlice)

	// Print the slice again in the main function
	fmt.Println("Original slice after modification:", originalSlice)
}
