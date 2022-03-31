package main

import (
	"os"

	"github.com/01-edu/z01"
)

/*for a := 0; a < lenght; a++ {
	for b := 0; b < lenght; b++ {
		if number[a] < number[b] {
			number[b], number[a] = number[a], number[b]
		}
	}
}*/

func main() {
	args := os.Args // stock argument dans args
	var longueur int
	for l := range args { // longeur sera égale au nombre d'arguments
		longueur = l
	}

	for i := longueur; i > 0; i-- { // temps quon est pas arrivé au debut de la chaine
		for _, j := range args[i] { // Affiche chaque argument
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
