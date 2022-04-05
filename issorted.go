package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	res1 := 1
	res2 := 1
	res3 := 1

	for index, value := range a {
		if index != len(a)-1 {
			if f(value, a[index+1]) < 0 {
				res1++
			} else if f(value, a[index+1]) > 0 {
				res2++
			} else if f(value, a[index+1]) == 0 {
				res3++
			}
		}
	}
	return res1 == len(a) || res2 == len(a) || res3 == len(a) // retourne true si une des conditions marche sinon retourne false
}
