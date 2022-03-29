package piscine

func NRune(s string, n int) rune {
	r := []rune(s)
	if n > 0 && n <= StrLen(s) {
		for i := 0; StrLen(s) != n+1; i++ {
			if i == n+1 {
				return r[n-1]
			}
		}
	}
	return 0
}
