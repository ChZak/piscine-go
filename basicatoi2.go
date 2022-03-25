package piscine

func BasicAtoi2(s string) int {
	Mystr := []rune(s)
	count := 0
	for i := range Mystr {
		if Mystr[i] < 48 || Mystr[i] > 57 {
			return 0
		} else {
			count = count*10 + int(Mystr[i]-48)
			i++
		}
	}
	return count
}
