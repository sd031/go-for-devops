package main

import (
	"fmt"
)

func main() {
	// Initialize a map
	fruits := map[string]int{"apple": 5, "banana": 7}

	// Check Length
	fmt.Println("Length of the map:", len(fruits))

	// Adding Elements
	fruits["orange"] = 10
	fmt.Println("Added 'orange':", fruits)

	// Retrieving Elements
	applePrice := fruits["apple"]
	fmt.Println("Price of apple:", applePrice)

	// Checking Existence
	price, exists := fruits["kiwi"]
	if exists {
		fmt.Println("Price of kiwi:", price)
	} else {
		fmt.Println("Kiwi does not exist in the map")
	}

	orangePrice, exists := fruits["orange"]
	if exists {
		fmt.Println("Price of orange:", orangePrice)
	} else {
		fmt.Println("orange does not exist in the map")
	}

	// Deleting Elements
	delete(fruits, "banana")
	fmt.Println("After deleting 'banana':", fruits)
}
