package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("example.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("\nNew content\n"); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File updated successfully.")
}
