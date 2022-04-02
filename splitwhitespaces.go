package piscine

func SplitWhiteSpaces(s string) []string {
	myStr := ""
	myTab := []string{}
	for i, char := range s {
		if i == len(s)-1 && string(char) != " " { // si il y a pas despace
			myStr += string(char)        // ajoute les charactere a ma nouvelle string
			myTab = append(myTab, myStr) // et ajoute la string a mon tableau
		} else if string(char) != " " && string(char) != "\n" { // sinon si on est pas a la fin de la chaine et qu'on a pas rencontrer d'espace
			myStr += string(char) // Ajoute les characetre a ma string
		} else {
			if len(myStr) >= 1 { // si il y a un ou pleusiur mot
				myTab = append(myTab, myStr) // ajouter les a mon tableau
			}
			myStr = "" // remet la chaine a zero pour chaque mot
		}
	}
	return myTab
}
