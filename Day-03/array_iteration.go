package main

import "fmt"

func main() {
	// Define an array of integers
	numbers := [5]int{10, 20, 30, 40, 50}

	// Using a traditional for loop to iterate over the array
	fmt.Println("Iterating using traditional for loop:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index: %d, Value: %d\n", i, numbers[i])
	}

	// Using a range loop to iterate over the array
	fmt.Println("\nIterating using range loop:")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
