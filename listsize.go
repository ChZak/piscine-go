package piscine

func ListSize(l *List) int {
	size := 0
	if l.Head.Next != nil {
		size++
	}
	return size
}
