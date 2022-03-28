package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i <= nb/2; i++ { // si il y a un chiffre apres la virgule lors de la division, c'est un nombre premier
		if nb%i == 0 { // si le chiffre apres la virgule est zero apres la division, ce n'est pas un nomnbre premier
			return false
		}
	}
	return true
}
