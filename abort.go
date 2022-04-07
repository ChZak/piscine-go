package piscine

func Abort(a, b, c, d, e int) int {
	args := []int{a, b, c, d, e}
	var myArray []int
	for i := 0; i < len(args); i++ {
		for j := 1; j < len(args)-i; j++ {
			if args[j] < args[j-1] {
				args[j], args[j-1] = args[j-1], args[j]
			}
		}
		myArray = append(myArray, args[i])
	}
	return myArray[2]
}
