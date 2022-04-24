package main

import "fmt"

func fibonacci(number int) int {
	var result = []int{0, 1}
	x := 0
	y := 1

	if number < 2 {
		return result[number]
	}

	for y < number {
		result = append(result, result[x]+result[y])
		x++
		y++
	}

	// fmt.Println(result)
	return result[y]
}

func main() {
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(9))
	fmt.Println(fibonacci(10))
	fmt.Println(fibonacci(12))
}
