package factorial

import "fmt"

func FactorialMain() {
	var number int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println("Factorial of", number, "is => ", factorial(number))
}

func factorial(n int) int {
	if n < 1 {
		return 1
	}
	return n * factorial(n-1)
}
