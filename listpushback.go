package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL // Noeud principal
	Tail *NodeL // Noeud final
}

func ListPushBack(l *List, data interface{}) {
	ele := &NodeL{ // Ajoute les data a Data
		Data: data,
	}
	if l.Head == nil { // Si le premier noeud de la liste est vide
		l.Head = ele // Il ajoute data a Head (premier element de la liste)
	} else {
		current := l.Head         // current contient le premier noeud de la liste
		for current.Next != nil { // si le noeud suivant n'est pas nil ca veut dire quil y a dautre noeud
			current = current.Next // Temps qu'on est pas arriver a la fin de la list continue
		}
		current.Next = ele // On ajoute 'ele' au noeud suivant
	}
}
