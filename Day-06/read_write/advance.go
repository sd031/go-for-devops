package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create or open the file
	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %s\n", err)
		return
	}
	defer file.Close()

	// Create a writer
	writer := bufio.NewWriter(file)

	// Create a scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text (type 'exit' to quit): ")

	// Read input line by line
	for scanner.Scan() {
		input := scanner.Text()

		// Check for the exit command
		if input == "exit" {
			break
		}

		// Write the input to the file
		if _, err := writer.WriteString(input + "\n"); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write to file: %s\n", err)
			return
		}

		// Flush the writer to ensure all data is written to the file
		writer.Flush()

		fmt.Print("Enter text (type 'exit' to quit): ")
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
}
