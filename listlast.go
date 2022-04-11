package piscine

func ListLast(l *List) interface{} {
	temp := l.Head
	for temp != nil {
		if temp.Next == nil {
			return temp.Data
		}
		temp = temp.Next
	}
	return temp
}
