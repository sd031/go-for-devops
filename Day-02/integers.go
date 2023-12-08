package main

import "fmt"

func main() {
	fmt.Println("Signed integers:")
	var a int8 = 127     // ranges from -128 to 127
	var b int16 = -32768 // ranges from -32768 to 32767
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	fmt.Println("Unsigned integers:")
	var c uint8 = 255    // ranges from 0 to 255
	var d uint16 = 65535 // ranges from 0 to 65535
	fmt.Println("c=", c)
	fmt.Println("d=", d)

	fmt.Println("Machine Dependent Types:")
	var e int = 123456789 // size depends on the architecture
	var f uint = 123456789
	var g uintptr = 0xdeadbeef
	fmt.Println("e=", e)
	fmt.Println("f=", f)
	fmt.Println("g=", g)

}
