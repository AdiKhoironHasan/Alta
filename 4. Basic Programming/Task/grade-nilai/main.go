package main

import "fmt"

func main() {
	var studentScore int = 80

	if studentScore >= 80 && studentScore <= 100 {
		fmt.Println("Nilai A")
	} else if studentScore >= 65 && studentScore <= 79 {
		fmt.Println("Nilai B")
	} else if studentScore >= 50 && studentScore <= 64 {
		fmt.Println("Nilai C")
	} else if studentScore >= 35 && studentScore <= 49 {
		fmt.Println("Nilai D")
	} else if studentScore >= 0 && studentScore <= 34 {
		fmt.Println("Nilai E")
	} else {
		fmt.Println("Nilai Invalid")
	}
}
