package main

import "fmt"

func playWithAsterik(n int) {
	var i, j int
	for i = 1; i <= n; i++ {
		for j = 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k != (2*i - 1); k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func main() {
	playWithAsterik(5)
}
