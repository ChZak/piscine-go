package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, str := range a { // parcour les chaines du tableau a
		for _, char := range []rune(str) { // // parcour les charactere de la chaine
			z01.PrintRune(char) // affiche les charactere
		}
		z01.PrintRune('\n')
	}
}
