package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x string
	y string
}

// Affiche une string
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func setPoint(ptr *point) {
	ptr.x = "42"
	ptr.y = "21"
}

func main() {
	points := &point{} // Pointeur vers la strucutre point
	setPoint(points)
	xstr := "x = "
	ystr := ", y = "
	printStr(xstr)
	printStr(points.x)
	printStr(ystr)
	printStr(points.y)
	z01.PrintRune('\n')
}
