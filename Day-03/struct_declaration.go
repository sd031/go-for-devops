package main

import "fmt"

func main() {
	// Basic declaration of a struct
	type Person struct {
		Name string
		Age  int
	}

	// Declaring and Instantiating together
	john := Person{Name: "John", Age: 30}
	fmt.Println("Declared and Instantiated together:", john)

	// Using the new keyword
	jane := new(Person)
	jane.Name = "Jane"
	jane.Age = 25
	fmt.Println("Using new keyword:", *jane)

	// Anonymous structs
	anon := struct {
		Country string
		Code    int
	}{
		Country: "USA",
		Code:    1,
	}
	fmt.Println("Anonymous struct:", anon)

	// Nested structs
	type Address struct {
		City  string
		State string
	}

	type Employee struct {
		PersonalInfo Person
		Address      Address
	}

	emp := Employee{
		PersonalInfo: Person{Name: "Alice", Age: 28},
		Address:      Address{City: "New York", State: "NY"},
	}
	fmt.Println("Nested structs:", emp)

	// Embedded (Anonymous) Fields
	type Manager struct {
		Person
		Department string
	}

	mgr := Manager{
		Person:     Person{Name: "Bob", Age: 35},
		Department: "Finance",
	}
	fmt.Println("Embedded fields:", mgr)
}
