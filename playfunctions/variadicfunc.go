package playfunctions

import "fmt"

func VariadicFunc() {
	// Example of variadic function using multiple inputs to the function
	fmt.Println("Sum of numbers 1-10 is => ", sumInts(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func sumInts(num ...int64) int64 { // Usage of variadic function declaration with multiple inputs
	var sum int64 = 0
	for _, v := range num {
		sum += v
	}
	return sum
}
