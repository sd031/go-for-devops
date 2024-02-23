package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{"Sandip", 32}

	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	var person2 Person

	newErr := json.Unmarshal(jsonData, &person2)
	if newErr != nil {
		log.Fatal(newErr)
	}

	fmt.Printf("%+v\n", person2)
}
