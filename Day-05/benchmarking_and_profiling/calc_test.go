package calc

import "testing"

// BenchmarkCalculateSum benchmarks the CalculateSum function
func BenchmarkCalculateSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSum(100)
	}
}
