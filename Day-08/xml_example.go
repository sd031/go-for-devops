package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

func main() {
	person := Person{Name: "Sandip", Age: 32}

	xmlData, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(xmlData))

	var person2 Person

	newErr := xml.Unmarshal(xmlData, &person2)
	if newErr != nil {
		log.Fatal(newErr)
	}

	fmt.Printf("%+v\n", person2)
}
