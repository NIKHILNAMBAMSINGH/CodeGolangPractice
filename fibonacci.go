package main

import "fmt"

func fibonacci(value int) []int {

	firstTerm := 0
	secondTerm := 1
	slices := make([]int, 0)
	for i := 0; i <= value; i++ {
		slices = append(slices, firstTerm)
		fmt.Print(firstTerm)
		thirdtime := firstTerm + secondTerm
		firstTerm = secondTerm
		secondTerm = thirdtime

	}
	return slices

}
func main() {
	value := 5
	result := fibonacci(value)
	for i, val := range result {
		fmt.Println(i, val)
	}

}
