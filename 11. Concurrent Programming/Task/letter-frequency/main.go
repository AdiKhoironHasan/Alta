package main

import (
	"fmt"
	"time"
)

func main() {
	leters := "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Maiores illo dolores labore nemo, accusamus laudantium fugit reiciendis numquam id, suscipit sed. Expedita quasi deleniti provident culpa ut illo, sit suscipit."

	go func(leters string) {
		result := map[string]int{}
		for _, val := range leters {
			if string(val) == " " {
				continue
			}

			if _, ok := result[string(val)]; ok {
				result[string(val)] += 1
			} else {
				result[string(val)] = 1
			}
		}
		for key, val := range result {
			fmt.Println(key, ":", val)
		}
	}(leters)

	time.Sleep(1 * time.Millisecond)
}
