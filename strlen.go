package piscine

func StrLen(s string) int {
	Mystr := []rune(s)
	c := 0
	for i := range Mystr {
		i++
		c++
	}
	return c
}
