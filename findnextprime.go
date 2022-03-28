package piscine

import "piscine"

func FindNextPrime(nb int) int {
	if piscine.IsPrime(nb) == true {
		return nb
	} else {
		n := nb + 1
		for piscine.IsPrime(n) == false {
			n++
		}
		return n
	}
}
