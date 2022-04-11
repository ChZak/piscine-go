package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	temp := l
	size := 0
	for temp != nil && size < pos {
		temp = temp.Next
		size++
	}
	return temp
}
