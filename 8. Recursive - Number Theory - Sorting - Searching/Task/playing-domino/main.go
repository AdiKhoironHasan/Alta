package main

import (
	"fmt"
)

func playingDomino(cards [][]int, deck []int) interface{} {
	var result = [2]int{}
	var list = [][]int{}

	for _, val := range cards {
		if val[0] == deck[0] || val[0] == deck[1] || val[1] == deck[0] || val[1] == deck[1] {
			list = append(list, val)
		}
	}

	if len(list) <= 0 {
		return "tutup kartu"
	}

	result[0] = list[0][0]
	result[1] = list[0][1]
	for _, val := range list {
		if val[0] > result[0] || val[0] > result[1] || val[1] > result[0] || val[1] > result[1] {
			result[0] = val[0]
			result[1] = val[1]
		}
	}

	return result
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
}
