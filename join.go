package piscine

func Join(strs []string, sep string) string {
	var res string
	size := 0
	for i := range strs {
		i++
		size++
	}
	for index, elements := range strs {
		res += elements
		if index != size-1 {
			res += sep
		}
	}
	return res
}
