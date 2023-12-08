package main

import "fmt"

func main() {
	fmt.Println("map:")
	r := map[string]int{"Alice": 30, "Bob": 25} // map with string keys and int values
	fmt.Println("r=", r)
	fmt.Println("Alice=", r["Alice"])
	fmt.Println("Bob=", r["Bob"])
}
