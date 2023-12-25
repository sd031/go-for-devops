package main

import "fmt"

func main() {
	// 1. Declaration Without Initialization
	var arr1 [5]int
	fmt.Println("Array declared without initialization:", arr1)

	// 2. Declaration With Initialization
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array declared with initialization:", arr2)

	// 3. Short Hand Declaration
	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array with short hand declaration:", arr3)

	// 4. Initializing With Specific Elements
	arr4 := [5]int{1: 10, 3: 30}
	fmt.Println("Array initializing with specific elements:", arr4)

	// 5. Using the ... Operator to Infer Length
	arr5 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Array using ... operator to infer length:", arr5)

	// 6. Array of Arrays (Multi-dimensional Arrays)
	var multiArray [2][3]int
	multiArray[0] = [3]int{1, 2, 3}
	multiArray[1] = [3]int{4, 5, 6}
	fmt.Println("Multi-dimensional Array:", multiArray)

	// 7. Initializing Multi-dimensional Arrays
	initializedMultiArray := [2][3]int{{7, 8, 9}, {10, 11, 12}}
	fmt.Println("Initialized Multi-dimensional Array:", initializedMultiArray)

	// 8. Arrays with Pointers
	var pointerArray [3]*int
	num1, num2, num3 := 13, 14, 15
	pointerArray[0] = &num1
	pointerArray[1] = &num2
	pointerArray[2] = &num3
	fmt.Println("Array with Pointers:")
	for i := 0; i < len(pointerArray); i++ {
		fmt.Printf("pointerArray[%d] = %d\n", i, *pointerArray[i])
	}

	// 9. Arrays of Structs
	type Person struct {
		Name string
		Age  int
	}
	var structArray [2]Person
	structArray[0] = Person{"Alice", 30}
	structArray[1] = Person{"Bob", 25}
	fmt.Println("Array of Structs:", structArray)
}
