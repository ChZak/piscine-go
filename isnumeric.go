package piscine

func IsNumeric(s string) bool {
	str := []rune(s)
	for r := 0; r < len(str); r++ {
		if str[r] < '0' || str[r] > '9' {
			return false
		}
	}
	return true
}
