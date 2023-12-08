package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Basic switch:")
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	// ...
	default:
		fmt.Println("Invalid day")
	}

	fmt.Println("switch with Initialization:")
	switch today := time.Now().Weekday(); today {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	fmt.Println("switch Without a Condition (like if-else):")
	x := 42
	switch {
	case x < 0:
		fmt.Println("x is negative")
	case x == 0:
		fmt.Println("x is zero")
	default:
		fmt.Println("x is positive")
	}

}
