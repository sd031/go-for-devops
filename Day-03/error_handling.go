package main

import (
	"errors"
	"fmt"
)

// CustomError is a custom error type
type CustomError struct {
	msg string
}

// Error implements the error interface for CustomError
func (e *CustomError) Error() string {
	return e.msg
}

// SomeFunction demonstrates returning a custom error
func SomeFunction(flag bool) error {
	if !flag {
		// Return a custom error
		return &CustomError{"Custom error occurred"}
	}
	return nil
}

// SafeFunction demonstrates using panic and recover
func SafeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// This would cause a panic
	panic("A problem occurred")
}

// divide performs division and handles division by zero error.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	// Simple error return
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Another example with valid division
	result, err = divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	// Example of checking and handling errors with custom errors:
	err2 := SomeFunction(false)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}

	// Example of using panic and recover
	SafeFunction()
}
