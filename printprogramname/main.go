package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := []rune(os.Args[0])
	for i := range name {
		z01.PrintRune(name[i])
	}
}
