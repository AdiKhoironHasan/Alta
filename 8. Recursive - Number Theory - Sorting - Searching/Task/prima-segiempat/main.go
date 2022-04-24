package main

import "fmt"

func primaSegiEmpat(high, wide, start int) {
	var prime = []int{}
	var sum int
	for i := start + 1; i > -1; i++ {
		if i == 1 {
			continue
		}

		if len(prime) >= high*wide {
			break
		}

		if primeNumber(i) == true {
			prime = append(prime, i)
			sum += i
		}
	}

	count := 0
	for i := 0; i < wide; i++ {
		for y := 0; y < high; y++ {
			fmt.Print(prime[count], " ")
			count += 1
		}

		fmt.Print("\n")
	}

	fmt.Print(sum, "\n\n")
}

func primeNumber(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		} else {
			continue
		}
	}

	return true
}

func main() {
	primaSegiEmpat(2, 3, 13)
	primaSegiEmpat(5, 2, 1)

}
