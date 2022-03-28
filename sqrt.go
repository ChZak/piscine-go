package piscine

func Sqrt(nb int) int {
	if nb > 0 {
		res := 1
		sqrt := 0
		for i := 1; res <= nb; i++ {
			res = i * i
			sqrt++
			if res == nb {
				return sqrt
			}
		}
		return 0
	} else {
		return 0
	}
}
