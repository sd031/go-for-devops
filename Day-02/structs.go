package main

import "fmt"

func main() {
	fmt.Println("Structs:")
	type Person struct {
		Name string
		Age  int
	}
	var q Person = Person{"Alice", 30}
	fmt.Println("q=", q)
	fmt.Println("name=", q.Name)
	fmt.Println("age=", q.Age)
}
