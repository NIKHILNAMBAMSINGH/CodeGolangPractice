package main

import (
	"fmt"
)

func reverseArray(input []int) []int {
	length := len(input)
	reversedArray := make([]int, length)
	for i := 0; i < length; i++ {
		reversedArray[length-i-1] = input[i]
	}
	fmt.Println(reversedArray)
	return reversedArray
}

func sameArrayOrNot(a, o []int) bool {
	if len(a) != len(o) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != o[i] {
			return false
		}

	}
	return true
}

func main() {
	array := make([]int, 0)
	var n int
	for i := 0; i < 3; i++ {
		fmt.Scan(&n)
		array = append(array, n)
	}
	originalArray := array
	reverseArray := reverseArray(array)

	result := sameArrayOrNot(originalArray, reverseArray)
	if result {
		fmt.Println("Same")
	} else {
		fmt.Println("Not same")
	}

}
