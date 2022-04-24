package main

import (
	"fmt"
)

func pow(x, n int) int {
	result := x
	for n > 1 {
		n--
		result *= x
	}

	return result
}

func main() {
	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
	fmt.Println(pow(10, 2))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(7, 3))
}
