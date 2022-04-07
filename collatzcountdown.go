package piscine

func CollatzCountdown(start int) int {
	result := 0
	if start <= 0 {
		return -1
	} else if start%2 == 0 {
		result += start / 2
	} else {
		result += start*3 + 1
	}
	return result
}
