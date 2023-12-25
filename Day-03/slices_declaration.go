package main

import "fmt"

func main() {
	// 1. Using the var Keyword without Initializing
	var slice1 []int
	fmt.Println("Slice 1 (uninitialized):", slice1)

	// Initialize slice1 before using it
	slice1 = []int{1, 2, 3}
	fmt.Println("Slice 1 (initialized):", slice1)

	// 2. Using a Slice Literal
	slice2 := []int{4, 5, 6}
	fmt.Println("Slice 2:", slice2)

	// 3. Using the make Function
	slice3 := make([]int, 3) // Length and capacity are 3
	fmt.Println("Slice 3 before putting value:", slice3)
	slice3[0] = 7
	slice3[1] = 8
	slice3[2] = 9
	fmt.Println("Slice 3:", slice3)
	slice4 := make([]int, 3, 5) // Length 3 and capacity are 3
	slice4[0] = 7
	slice4[1] = 8
	slice4[2] = 9
	fmt.Println("Slice 4:", slice4)
	slice4 = append(slice4, 10)
	slice4 = append(slice4, 11)
	fmt.Println("Slice 4 after append:", slice4)
	fmt.Println("Slice 4 Capacity:", cap(slice4))
	slice4 = append(slice4, 12)
	slice4 = append(slice4, 13)
	fmt.Println("Slice 4 after append again:", slice4)
	fmt.Println("Slice 4 Capcity:", cap(slice4))

	// 4. From an Array
	array := [5]int{10, 11, 12, 13, 14}
	slice5 := array[1:4] // Slicing from index 1 to 3 (excluding index 4)
	fmt.Println("Slice 5:", slice5)
}
