package piscine

func Map(f func(int) bool, a []int) []bool {
	array := make([]bool, len(a))
	for index, value := range a {
		array[index] = f(value)
	}
	return array
}
