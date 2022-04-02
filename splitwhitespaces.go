package piscine

import (
	"strings"
)

func SplitWhiteSpaces(s string) []string {
	myArray := []string(strings.Fields(s))
	return myArray
}
