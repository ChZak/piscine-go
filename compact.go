package piscine

func Compact(ptr *[]string) int {
	count := 0
	for _, i := range *ptr {
		if i != "" {
			count++
		}
	}
	newArray := make([]string, count)
	n := 0
	for _, str := range *ptr {
		if str != "" {
			*ptr = newArray   // ajout a ptr, ajoute rien au debut ce qui rend la chaine vide
			newArray[n] = str // ajoute de la chaine a new array
			n++
		}
	}
	return count
}
