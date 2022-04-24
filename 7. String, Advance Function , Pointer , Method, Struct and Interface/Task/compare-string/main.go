package main

import (
	"fmt"
)

func Compare(a, b string) string {
	// var y string
	var n string

	if len(a) < len(b) {
		for key, _ := range a {
			n = b[:key+1]
		}
	} else {
		for key, _ := range b {
			n = a[:key+1]
		}
	}

	return n
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU_KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
