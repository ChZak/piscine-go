package piscine

func Rot14(s string) string {
	myStr := ""
	for _, char := range s {
		if char >= 'a' && char <= 'l' || char >= 'A' && char <= 'L' {
			myStr += string(char + 14)
		} else if char >= 'm' && char <= 'z' || char >= 'M' && char <= 'Z' {
			myStr += string(char - 12)
		} else {
			myStr += string(char)
		}
	}
	return myStr
}
