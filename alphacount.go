package piscine

func AlphaCount(s string) int {
	myS := []rune(s)
	count := 0
	for i := 0; i < StrLen(s); i++ {
		if myS[i] >= 41 && myS[i] <= 90 {
			count += 1
		} else if myS[i] >= 97 && myS[i] <= 122 {
			count += 1
		} else {
			count -= 1
		}
	}
	return count + 1
}
