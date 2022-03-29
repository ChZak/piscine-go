package piscine

func Index(s string, toFind string) int {
	length := len(s)
	sublength := len(toFind)

	t := 0
	for i := 0; i < length; i++ {
		j := 0
		t = i
		for j < sublength {
			if t < length && s[t] == toFind[j] {
				j++
				t++
			} else {
				break
			}
		}
		if j == sublength {
			return i
		}
	}
	return -1
}
