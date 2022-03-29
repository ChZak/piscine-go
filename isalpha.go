package piscine

func IsAlpha(s string) bool {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		if str[i] < 33 || str[i] > 126 {
			return false
		}
	}
	return true
}
