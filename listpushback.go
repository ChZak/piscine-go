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
	ele := &NodeL{
		Data: data,
	}
	if l.Head == nil {
		l.Head = ele
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = ele
	}
}
