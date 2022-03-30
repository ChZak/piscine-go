package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var number []int

	if n > 0 {
		for i := n; i > 0; i = i / 10 { // la division par 10 permet de recuperer le chiffre le plus petit via le modulo
			number = append(number, i%10)
		}

		for _, letter := range number { // affiche chaque chiffre
			z01.PrintRune(rune(letter + '0')) // le 0 permet de commencer des le debut
		}
	} else {
		z01.PrintRune('0')
	}
}
