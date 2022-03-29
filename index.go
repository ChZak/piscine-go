package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(toFind); j++ {
			if s[i] == toFind[j] {
				if Compare(s, toFind)+1 == 0 {
					return i
				}
			}
		}
	}
	return -1
}
