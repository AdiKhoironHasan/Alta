package main

import "fmt"

// with Tabbulation (Bottom Up)
// penggunaan memory lebih banyak
// cpu sedikit
func fibo(n int) int{
	var memo = map[int]int{}

	for i:=0; i<=n;i++{
		if i<= 1{
			memo[i] = i
		}else{
			memo[i] = memo[i-1] + memo[i-2]
		}
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