package main

import (
	"fmt"
	"sort"
)

func MaximumBuyProduct(money int, productPrice []int) {
	var result int
	sort.Ints(productPrice)

	for _, val := range productPrice {
		money -= val
		if money >= 0 {
			result += 1
		}
	}

	fmt.Println(result)
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 1000, 14000})
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})
	MaximumBuyProduct(0, []int{1000, 30000})

}
