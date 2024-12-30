package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := transformNumbers(&numbers, double)
	fmt.Println("Doubled: ", doubled)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println("Tripled: ", tripled)
	squared := transformNumbers(&numbers, square)
	fmt.Println("Squared: ", squared)
}

func transformNumbers(n *[]int, transform func(int) int) []int {
	dNum := []int{}
	for _, v := range *n {
		dNum = append(dNum, transform(v))
	}
	return dNum
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}

func square(num int) int {
	return num * num
}
