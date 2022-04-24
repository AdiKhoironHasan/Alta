package main

import "fmt"

type pair struct {
	name string
	coun int
}

func MostAppearItem(items []string) []pair {
	datas := map[string]int{}
	var result []pair
	for _, val := range items {
		datas[val] += 1
	}

	for key, val := range datas {
		data := pair{
			name: key,
			coun: val,
		}
		result = append(result, data)
	}

	return result
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"FOOTBALL", "BASKETBALL", "TENIS"}))
}
