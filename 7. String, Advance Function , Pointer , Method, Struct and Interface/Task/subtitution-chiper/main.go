package main

import (
	"fmt"
	"strings"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""

	var builder strings.Builder
	value := strings.ToUpper(s.name)

	for _, r := range value {
		if r >= 'A' && r <= 'Z' {
			r = caesarCipherShiftRune(r, uint(s.score))
		}

		builder.WriteRune(r)
	}

	nameEncode = builder.String()
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	nameDecode = s.Encode()
	return nameDecode
}

func caesarCipherShiftRune(r rune, shift uint) rune {
	s := rune(shift % 26)

	if s == 0 {
		return r
	}

	return (((r - 'A') + s) % 26) + 'A'
}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a

	fmt.Println("[1] Encrypt \n[2] Decrypt \nChoose your menu!")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name " + a.name + " is: " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name " + a.nameEncode + " is: " + c.Decode())
	} else {
		fmt.Println("Wrong innput name menu")
	}
}
