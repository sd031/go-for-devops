package main

import (
	"fmt"
)

// Function to modify an element of the array
func modifyArray(arr [5]int, index int, value int) [5]int {
	arr[index] = value
	return arr
}

func main() {
	// Initialize an array
	originalArray := [5]int{1, 2, 3, 4, 5}

	// Print the length of the array
	fmt.Println("Length of the array is:", len(originalArray))

	// Copy the array and print the copied array
	var copiedArray = originalArray // This creates a copy of myArray
	fmt.Println("Copied Array:", copiedArray)

	// Modify the third element (index 2) of the original array to 10
	modifiedArray := modifyArray(originalArray, 2, 10)
	fmt.Println("Modified Array:", modifiedArray)

	// Compare the original and modified arrays
	isEqual := originalArray == modifiedArray // isEqual is false
	fmt.Println("Are original and modified arrays equal?", isEqual)

	isCopiedEqual := originalArray == copiedArray // isCopiedEqual is true
	fmt.Println("Are original and copied arrays equal?", isCopiedEqual)
}
