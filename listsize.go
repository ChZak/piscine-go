package piscine

func ListSize(l *List) int {
	size := 0
	for l.Head.Next != nil {
		size++
	}
	return size
}
