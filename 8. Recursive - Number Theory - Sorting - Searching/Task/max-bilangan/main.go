package main

import "fmt"

func MaxSequence(arr []int) int {
	var result int
	for _, val := range arr {
		result += val
	}
	return result
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
