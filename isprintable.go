package piscine

func IsPrintable(s string) bool {
	str := []rune(s)
	for r := 0; r < len(str); r++ {
		if str[r] < 33 || str[r] > 126 {
			return false
		}
	}
	return true
}
