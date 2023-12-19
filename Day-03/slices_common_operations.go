package main

import (
	"fmt"
	"sort"
)

func main() {
	// Initialize a slice
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	fmt.Println("Original Slice", slice)

	// Print Length and Capacity of the slice
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))

	// Appending to a Slice
	slice = append(slice, 7)
	fmt.Println("After appending:", slice)

	// Slicing a Slice
	subSlice := slice[2:5]
	fmt.Println("Sub-slice:", subSlice)

	// Delete an element from slice (e.g., at index 3)
	slice = append(slice[:3], slice[4:]...)
	fmt.Println("After deleting an element:", slice)

	// Checking if a Slice is Empty
	if len(slice) == 0 {
		fmt.Println("Slice is empty")
	} else {
		fmt.Println("Slice is not empty")
	}

	// Sorting a Slice
	sort.Ints(slice)
	fmt.Println("Sorted slice:", slice)

	// Finding an Element in a Slice
	elementToFind := 5
	found := false
	for _, v := range slice {
		if v == elementToFind {
			found = true
			break
		}
	}
	if found {
		fmt.Printf("Element %d found in the slice\n", elementToFind)
	} else {
		fmt.Printf("Element %d not found in the slice\n", elementToFind)
	}
}
