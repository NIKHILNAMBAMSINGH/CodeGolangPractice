package factorial

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

}

func main() {
	a := 5
	value := factorial(a)
	fmt.Println(value)
}
