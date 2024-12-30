package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := doubleNumbers(&numbers)
	fmt.Println(doubled)
}

func doubleNumbers(n *[]int) []int {
	dNum := []int{}
	for _, v := range *n {
		dNum = append(dNum, v*2)
	}
	return dNum
}
