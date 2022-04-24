package main

import "fmt"

// with memorization (Top Down)
var memo = map[int]int{}
func fibo(n int) int {
	key, val := memo[n]
	if val {
		return key
	}

	if n <= 1{
		memo[n] = n
	}else{
		memo[n] = fibo(n-1) + fibo(n-2)
	}

	return memo[n]
}

func main(){
	fmt.Println(fibo(0))
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(3))
	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))
	fmt.Println(fibo(9))
	fmt.Println(fibo(10))
}