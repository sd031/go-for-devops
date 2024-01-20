package main

import (
	"fmt"
	"os"
)

func main() {
	// Setting an environment variable
	os.Setenv("MY_ENV_VAR", "Hello World")

	// Reading an environment variable
	myVar := os.Getenv("MY_ENV_VAR")
	if myVar == "" {
		fmt.Println("MY_ENV_VAR is not set")
	} else {
		fmt.Printf("MY_ENV_VAR: %s\n", myVar)
	}

	// Listing all environment variables
	// fmt.Println("All Environment Variables:")
	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }
}
