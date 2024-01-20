package calc

// CalculateSum calculates the sum of all numbers up to n
func CalculateSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
