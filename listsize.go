package piscine

func ListSize(l *List) int {
	size := 0
	temp := l.Head
	for temp != nil {
		temp = temp.Next
		size++
	}
	return size
}
