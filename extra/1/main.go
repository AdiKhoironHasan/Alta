package main

import (
	"fmt"
	"sort"
)

func a(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func main() {
	fmt.Println(a([]int{1, 3, 4, 2, 3, 4}))
}
