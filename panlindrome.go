package main

import (
	"fmt"
	"strings"
)

func palindrome(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	fmt.Println(s)
	length := len(s)
	newString := make([]rune, length)
	for i := 0; i < length; i++ {
		newString[length-i-1] = rune(s[i])
	}
	return string(newString)
}

func checking(s string) bool {
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[length-i-1] != s[i] {
			return false
		}
	}
	return true
}

func main() {
	s := "dog"
	reversedString := palindrome(s)
	fmt.Println(checking(reversedString))

}
