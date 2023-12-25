package main

import (
	"fmt"
)

// Implementing an Interface
type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Hello, " + p.Name
}

// Type Assertions
func PrintDetails(i interface{}) {
	str, ok := i.(string)
	if ok {
		fmt.Println("It's a string:", str)
	} else {
		fmt.Println("Not a string")
	}
}

// Interface with Type Switches
func Describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	// Implementing an Interface
	person := Person{Name: "Alice"}
	var greeter Greeter = person
	fmt.Println(greeter.Greet())

	// Type Assertions
	PrintDetails("Hello")
	PrintDetails(42)

	// Interface with Type Switches
	Describe("Hello, world!")
	Describe(2023)
	Describe(3.14)
}
