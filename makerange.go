package piscine

func MakeRange(min, max int) []int {
	var value []int
	if min >= max {
		return nil
	}
	value = make([]int, max-min)   // recupere le nombre de int a integrer dans le slice
	for i := 0; i < max-min; i++ { // du moment que l'on est pas arriver au nombre maximum
		value[i] = i + min // on ajoute la valuer dnas le slice value
	}
	return value
}
