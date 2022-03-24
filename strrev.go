package piscine

func StrRev(s string) string {
	rns := []rune(s)
	count := 0
	for i := range rns {
		i++
		count++
	}
	for i, j := 0, count-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
