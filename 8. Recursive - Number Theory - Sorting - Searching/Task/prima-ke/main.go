package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {
	var result int

	for i := 2; i > 0; i++ {

		check := checkPrime(i)
		if check == true {
			result += 1
		}

		if result >= number {
			result = i
			break
		}
	}

	return result
}

func checkPrime(n int) bool {
	sqrt := math.Sqrt(float64(n))
	sqrtNum := int(sqrt)

	for i := 2; i <= sqrtNum; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
