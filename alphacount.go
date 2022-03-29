package piscine

func AlphaCount(s string) int {
	myS := []rune(s)
	count := 1
	for i := 0; i < StrLen(s); i++ {
		if myS[i] >= 65 && myS[i] <= 90 {
			count += 1
		} else if myS[i] >= 97 && myS[i] <= 122 {
			count += 1
		}
	}
	return count - 1
}
