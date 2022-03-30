package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var number []int
	lenght := 0

	if n > 0 {
		for i := n; i > 0; i = i / 10 { // la division par 10 permet de recuperer le chiffre le plus petit via le modulo
			number = append(number, i%10)
		}
		for range number {
			lenght++
		}

		for a := 0; a < lenght; a++ {
			for b := 0; b < lenght; b++ {
				if number[a] < number[b] {
					number[b], number[a] = number[a], number[b]
				}
			}
		}

		for _, letter := range number { // affiche chaque chiffre
			z01.PrintRune(rune(letter + '0')) // le 0 permet de commencer des le debut
		}
	} else {
		z01.PrintRune('0')
	}
}
