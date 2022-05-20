package main

import "fmt"

func SimpleEquations(a, b, c int) {
	a = a + b + c
	b = a * b * c
	c = (a * a) + (b * b) + (c * c)

	if a >= 1 || c <= 1000 {
		fmt.Println(a, b, c)
	} else {
		fmt.Println("no solution")
	}
}

func main() {
	SimpleEquations(1, 2, 3)
	SimpleEquations(6, 6, 14)
}
