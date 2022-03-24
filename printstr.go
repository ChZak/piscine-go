package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, res := range []rune(s) {
		z01.PrintRune(res)
	}
}
