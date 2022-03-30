package piscine

func atoi(runes []rune) int {
	LenRune := 0
	result := 0
	for i := range runes {
		i++
		LenRune++
	}
	if LenRune == 0 {
		return 0
	}

	tens := 1
	for k := 0; k < LenRune-1; k++ {
		if runes[k] == '+' || runes[k] == '-' {
			continue
		}
		tens *= 10
	}

	for i := range runes {
		if (runes[0] == '-' || runes[0] == '+') && i == 0 {
			continue
		}
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}
		numb := 0
		for j := '0'; j < runes[i]; j++ {
			numb++
		}
		result += numb * tens
		tens /= 10
	}
	if runes[0] == '-' {
		result *= -1
	}
	return result
}

func TrimAtoi(s string) int {
	str := []rune(s)
	var number []rune
	j := 0

	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			number = append(number, str[i])
			j = i + 1
		} else if str[i] == '-' && j == 0 {
			number = append(number, '-')
		}
	}
	return atoi(number)
}
