package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	currentNode := l.Head
	for currentNode != nil {
		currentNode = currentNode.Next
		if comp(ref, currentNode.Data) == true {
			return &currentNode.Data
		}
	}
	return nil
}
