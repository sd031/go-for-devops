package main

import (
	"fmt"
)

func main() {
	// Define a 3D array
	// Here, we are creating a 3x3x3 array
	var array3D [3][3][3]int

	// Populate the 3D array with values
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				array3D[i][j][k] = i + j + k
			}
		}
	}

	// Print the 3D array
	for i := 0; i < 3; i++ {
		fmt.Printf("Array at i=%d:\n", i)
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				fmt.Printf("%d ", array3D[i][j][k])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
