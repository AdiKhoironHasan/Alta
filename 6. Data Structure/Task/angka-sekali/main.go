package main

import "fmt"

func munculSekali(angka string) []string {
	result := []string{}
	maps := map[string]bool{}
	for _, v := range angka {
		val := fmt.Sprintf("%c", v)

		if len(result) == 0 {
			result = append(result, val)
			maps[val] = true
		} else {
			if _, ok := maps[val]; !ok {
				maps[val] = true
			} else {
				maps[val] = false
			}
		}
	}

	finalResult := []string{}

	for k, e := range maps {
		if e == true {
			finalResult = append(finalResult, k)
		}
	}

	// fmt.Println(maps)

	return finalResult
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
