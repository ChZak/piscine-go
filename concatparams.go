package piscine

func ConcatParams(args []string) string {
	var res string
	size := 0
	for i := range args { // parcour les arguments, et defini la taille de chacun
		i++
		size++
	}
	for index, elements := range args {
		res += elements      // Ajoute chaque arguments dnas une string
		if index != size-1 { // si mon index est plus haut que la taille de mon argmuent, ça veut dire que l'on est a la fin d'un argmuent
			res += "\n" // Ajoute un retour a la ligne et passe au prochain arguments
		}
	}
	return res
}
