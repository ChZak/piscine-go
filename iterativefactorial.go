package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		res := 1
		for i := 1; i <= nb; i++ {
			res *= i
			if res > 9223372036854775807 {
				res = 0
				break
			}
		}
		return res
	}
}
