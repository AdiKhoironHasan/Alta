package main

import "fmt"

func Swap(a, b *int) {
	n := *a
	*a = *b
	*b = n
}

func main() {
	a := 10
	b := 20

	Swap(&a, &b)
	fmt.Println(a, b)
}
