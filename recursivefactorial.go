package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 1 || nb == 0 {
		return 1
	} else if nb >= 2147483647 {
		return 0
	}
	res := nb * RecursiveFactorial(nb-1)
	return res
}
