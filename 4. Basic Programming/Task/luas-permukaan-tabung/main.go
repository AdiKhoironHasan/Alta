package main

import "fmt"

func main() {
	const phi float64 = 3.14
	var r float64
	var T float64

	fmt.Scanf("%f\n", &r)
	fmt.Scanf("%f\n", &T)

	Lp := 2 * phi * r * (r + T)

	fmt.Println("luas permukaan adalah: ", Lp)
}
