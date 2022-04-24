package main

import "fmt"

// with memorization (Top Down)
// var memo = map[int]int{}
// func fibo(n int) int {
// 	key, val := memo[n]
// 	if val {
// 		return key
// 	}

// 	if n <= 1{
// 		memo[n] = n
// 	}else{
// 		memo[n] = fibo(n-1) + fibo(n-2)
// 	}

// 	return memo[n]
// }

// with Tabbulation (Bottom Up)
// penggunaan memory lebih banyak
// cpu sedikit
// func fibo(n int) int{
// 	var memo = map[int]int{}

// 	for i:=0; i<=n;i++{
// 		if i<= 1{
// 			memo[i] = i
// 		}else{
// 			memo[i] = memo[i-1] + memo[i-2]
// 		}
// 	}

// 	return memo[n]
// }

func fibo(n int) int {
	if n<=1{
		return n
	}

	// 0 0 1 1 2
	fiboMin1, fiboMin2, fiboI := 1,0,0
	for i:=2; i<=n; i++{
		fiboI = fiboMin1+fiboMin2
		fiboMin2 = fiboMin1
		fiboMin1 = fiboI
	}

	return fiboI
}

func main(){
	fmt.Println(fibo(50))
}