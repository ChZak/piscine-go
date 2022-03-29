package piscine

func LastRune(s string) rune {
	r := []rune(s)
	j := 0
	for i := range r {
		i++
		j++
	}
	return r[j-1]
}
