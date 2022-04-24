package main

import (
	"fmt"
	"sort"
)

func moneyCoins(money int) []int {
	coin := []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	var result = []int{}

	sort.Slice(coin, func(i, j int) bool {
		return coin[j] < coin[i]
	})

	for i := 0; i < len(coin); i++ {
		if (money - coin[i]) >= 0 {
			money -= coin[i]
			result = append(result, coin[i])
			i--
		}
	}

	return result
}

func main() {
	fmt.Println(moneyCoins(123))
	fmt.Println(moneyCoins(432))
	fmt.Println(moneyCoins(543))
	fmt.Println(moneyCoins(7752))
	fmt.Println(moneyCoins(15321))
}
