package main

import "fmt"

func main() {
	var y int     // 0
	var z bool    // false
	var aa string // ""
	var bb *int   // nil
	fmt.Println("Y=", y)
	fmt.Println("z=", z)
	fmt.Println("aa=", aa)
	fmt.Println("bb=", bb)
	// fmt.Println("bb=", bx) //enabling this line will through undefined error
}
