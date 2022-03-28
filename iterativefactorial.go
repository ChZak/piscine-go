package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 1 || nb == 0 {
		return 1
	} else {
		f := 1
		for i := 1; i < nb+1; i++ {
			f = f * i
			if f > 2147483647 {
				f = 0
				break
			}
		}
		return f
	}
}
