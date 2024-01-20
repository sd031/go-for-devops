package main

import (
	"fmt"
	"os"
)

func main() {
	// Checking if the right number of arguments are passed
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run simple_argument.go [filename] [message]")
		os.Exit(1)
	}

	// Getting command-line arguments
	filename := os.Args[1]
	message := os.Args[2]

	// Writing message to the specified file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(message + "\n")
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		return
	}

	fmt.Printf("Message '%s' written to file %s\n", message, filename)
}
