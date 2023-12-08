package main

import "fmt"

func main() {
	const Pi = 3.14
	// Pi = 400 // enbling this line will through error: cannot assign to Pi
	fmt.Println("Pi=", Pi)
	// fmt.Println("bb=", bx) //enabling this line will through undefined error
}
