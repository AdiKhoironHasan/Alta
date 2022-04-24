package main

import "fmt"

func cetakTabelPerkalian(number int) {
	for i := 1; i <= number; i++ {
		fmt.Printf("%d ", i)
		for k := 2; k <= number; k++ {
			fmt.Printf("%d ", i*k)
		}
		fmt.Printf("\n")
	}
}

func main() {
	cetakTabelPerkalian(9)
}
