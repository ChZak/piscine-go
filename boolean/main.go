package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Affiche une string
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// Vérifie si le chiffre est paire

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	args := os.Args[1:] // Recupère les argmuent sauf le nom du fichier
	length := 0

	// Compte le nombre d'arguments
	for range args {
		length++
	}

	if isEven(length) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
