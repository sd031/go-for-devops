package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Sample inputs to validate
	name := "John Doe"
	email := "john.doe@example.com"
	phoneNumber := "+1-234-567-8900"

	// Regex patterns
	namePattern := `^[a-zA-Z]+(?: [a-zA-Z]+)?$`
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phonePattern := `^\+\d{1,3}-\d{3}-\d{3}-\d{4}$`

	// Compile regex and match patterns
	matchName, _ := regexp.MatchString(namePattern, name)
	matchEmail, _ := regexp.MatchString(emailPattern, email)
	matchPhone, _ := regexp.MatchString(phonePattern, phoneNumber)

	// Output results
	fmt.Printf("Name Valid: %t\n", matchName)
	fmt.Printf("Email Valid: %t\n", matchEmail)
	fmt.Printf("Phone Number Valid: %t\n", matchPhone)
}
