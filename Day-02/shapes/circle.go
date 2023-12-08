// circle.go in the shapes package
package shapes

import "math"

// Exported Function
func AreaOfCircle(radius float64) float64 {
	return math.Pi * radius * radius
}

// unexported function
func diameterOfCircle(radius float64) float64 {
	return 2 * radius
}
