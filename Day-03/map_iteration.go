package main

import (
	"fmt"
)

func main() {
	// Initialize a map
	fruits := map[string]int{"apple": 5, "banana": 10}

	// Add an element
	fruits["orange"] = 7

	// Iterate over the map
	for key, value := range fruits {
		fmt.Printf("%s: %d\n", key, value)
	}
}
