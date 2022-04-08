package piscine

func ActiveBits(n int) int {
	count := 1
	for i := 0; i < n; i++ {
		count += n & 1
		n >>= 1
	}
	return count
}
