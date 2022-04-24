package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Avarage() float64 {
	var avg float64
	for _, v := range s.score {
		avg = (avg + float64(v))
	}
	fmt.Println(avg)
	avg = avg / float64(len(s.score))
	return avg
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	for k, v := range s.score {

		if v < min {
			min = v
			name = s.name[k]
		}

	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	for k, v := range s.score {

		if v > max {
			max = v
			name = s.name[k]
		}

	}
	return max, name
}

func main() {
	var a = Student{}

	for i := 0; i < 6; i++ {
		var name string

		fmt.Print("Input " + string(i) + " Student's Name: ")
		fmt.Scan(&name)
		a.name = append(a.name, name)

		var score int
		fmt.Print("Input " + name + " Score:")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAvarage Score Students is", a.Avarage())

	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is: "+nameMax+"(", scoreMax, ")")

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is: "+nameMin+"(", scoreMin, ")")
}
