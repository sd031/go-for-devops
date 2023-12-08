// square.go in the shapes package
package shapes

// Exported Function
func AreaOfSquare(side float64) float64 {
	return side * side
}

// unexported function
func perimeterOfSquare(side float64) float64 {
	return 4 * side
}
