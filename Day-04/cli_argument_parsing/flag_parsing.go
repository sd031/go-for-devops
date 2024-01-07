package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define flags
	filename := flag.String("file", "output.txt", "The name of the file to write to")
	message := flag.String("message", "Hello, World!", "The message to write in the file")

	// Parse the flags
	flag.Parse()

	// Open file
	file, err := os.OpenFile(*filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Write message to file
	_, err = file.WriteString(*message + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("Message written to %s\n", *filename)
}
