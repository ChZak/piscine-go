package piscine

func IsAlpha(s string) bool {
	str := []rune(s)
	for r := 0; r < len(str); r++ {
		if (str[r] < 'a' || str[r] > 'z') && (str[r] < 'A' || str[r] > 'Z') && (str[r] < '0' || str[r] > '9') {
			return false
		}
	}
	return true
}
