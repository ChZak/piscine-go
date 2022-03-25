package piscine

func BasicAtoi(s string) int {
	Mystr := []rune(s)
	count := 0
	for i := range Mystr {
		count = count*10 + int(Mystr[i]-48)
		i++
	}
	return count
}
