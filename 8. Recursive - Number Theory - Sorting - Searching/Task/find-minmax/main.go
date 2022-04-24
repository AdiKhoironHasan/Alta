package main

import "fmt"

func FindMinAndMax(arr []int) string {
	var max int = arr[0]
	var min int = arr[0]
	var maxString string = fmt.Sprintf("max: %d index: %d ", max, 0)
	var minString string = fmt.Sprintf("min: %d index: %d ", min, 0)

	for k, v := range arr {
		if v > max {
			max = v
			maxString = fmt.Sprintf("max: %d index: %d ", v, k)
		}

		if v < min {
			min = v
			minString = fmt.Sprintf("min: %d index: %d ", v, k)
		}
	}

	return minString + maxString
}

func main() {
	fmt.Println(FindMinAndMax([]int{5, 7, -2, -1, 8}))
	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	fmt.Println(FindMinAndMax([]int{4, 3, 9, -21, 7}))
	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))
}
