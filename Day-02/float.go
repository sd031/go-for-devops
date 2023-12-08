package main

import "fmt"

func main() {
	fmt.Println("float32 and float64:")
	var h float32 = 3.14              // single-precision floating-point
	var i float64 = 3.141592653589793 // double-precision
	fmt.Println("h=", h)
	fmt.Println("i=", i)

	fmt.Println("Complex Numbers:")
	var j complex64 = 1 + 2i  // complex number with float32 real and imaginary parts
	var k complex128 = 1 + 2i // complex number with float64 real and imaginary parts
	fmt.Println("j=", j)
	fmt.Println("k=", k)

}
