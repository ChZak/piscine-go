package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	r := []rune(s)
	for _, res := range r {
		z01.PrintRune(res)
	}
	z01.PrintRune('\n')
}
