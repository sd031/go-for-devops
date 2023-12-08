package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sum_of_num(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Basic Function:")
	result := add(3, 4)
	fmt.Println("result=", result)

	fmt.Println("Multiple Return Values:")
	a, b := swap("hello", "world")
	fmt.Println("a, b=", a, b)

	fmt.Println("Named Return Values:")
	sum := 100 // Example sum
	x, y := split(sum)
	fmt.Printf("The split of %d is: x = %d, y = %d\n", sum, x, y)

	fmt.Println("Variadic Functions:")
	sum_result := sum_of_num(1, 2, 3, 4)
	fmt.Println("The sum is:", sum_result)
}
