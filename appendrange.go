package piscine

func AppendRange(min, max int) []int {
	var value []int
	if min >= max {
		return nil
	} else {
		for i := min; i < max; i++ {
			value = append(value, i)
		}
	}
	return value
}
