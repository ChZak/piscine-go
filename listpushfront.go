package piscine

func ListPushFront(l *List, data interface{}) {
	ele := &NodeL{
		Data: data,
	}
	if l.Head == nil {
		l.Head = ele
	} else {
		ele.Next = l.Head // Ajoute le Head a l'élement suivant
		l.Head = ele      // Remplace head par lement ajouter
	}
}
