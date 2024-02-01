package piscine

func DescendAppendRange(max, min int) []int {
	var nb []int
	if max <= min {
		nb = []int{}
	}
	for i := max; i > min; i-- {
		nb = append(nb, i)
	}
	return nb
}
