package main

import (
	"day2/shapes" // replace with the actual path to the shapes package
	"fmt"
)

func main() {
	radius := 5.0
	areaCircle := shapes.AreaOfCircle(radius)
	fmt.Printf("Area of Circle: %f\n", areaCircle)

	side := 4.0
	areaSquare := shapes.AreaOfSquare(side)
	fmt.Printf("Area of Square: %f\n", areaSquare)

	// This will result in an error as diameterOfCircle is unexported
	// diameter := shapes.diameterOfCircle(radius)
}
