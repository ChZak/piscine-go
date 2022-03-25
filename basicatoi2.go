package main

import (
	"fmt"
)

func BasicAtoi2(s string) int {
	Mystr := []rune(s)
	count := 0
	for i := range Mystr {
		if Mystr[i] < 48 || Mystr[i] > 57 {
			return 0
		} else {
			count = count*10 + int(Mystr[i]-48)
			i++
		}
	}
	return count
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}
