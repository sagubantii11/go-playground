package playfunctions

import "fmt"

type fInt func(int) int

func PlayFunctions() {
	// Working with Anonymous Functions
	numbers := []int{1, 2, 3, 4, 5}
	doubled := transformNumbers(&numbers, func(num int) int { return num * 2 })
	fmt.Println("Doubled: ", doubled)
	tripled := transformNumbers(&numbers, func(num int) int { return num * 3 })
	fmt.Println("Tripled: ", tripled)
	squared := transformNumbers(&numbers, func(num int) int { return num * num })
	fmt.Println("Squared: ", squared)

	fmt.Println("Doubled again: ", transformNumbers(&doubled, createFunc(2))) // Usage of a function factory
}

func transformNumbers(num *[]int, transform fInt) []int {
	dNum := []int{}
	for _, v := range *num {
		dNum = append(dNum, transform(v))
	}
	return dNum
}

func createFunc(factor int) fInt {
	return func(num int) int {
		return num * factor
	}
}

// func double(num int) int {
// 	return num * 2
// }

// func triple(num int) int {
// 	return num * 3
// }

// func square(num int) int {
// 	return num * num
// }
