package main

import (
	"fmt"
	"os"
)

func main() {
	path := "example_folder"

	// Check if file/folder exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create folder if it doesn't exist
		err := os.Mkdir(path, 0755)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
		fmt.Println("Folder created successfully.")
	} else {
		fmt.Println("Folder already exists.")
	}
}
