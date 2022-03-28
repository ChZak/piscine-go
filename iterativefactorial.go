package piscine

func IterativeFactorial(nb int) int {
	f := 1
	if nb > 0 {
		for i := 1; i <= nb; i++ {
			f = f * i
		}
	} else {
		return 0
	}
	return f
}
