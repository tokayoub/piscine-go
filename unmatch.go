package piscine

func Unmatch(a []int) int {
	var i int
	for _, v := range a {
		i = 0
		for _, w := range a {
			if w == v {
				i++
			}
		}
		if i%2 != 0 {
			return v
		}
	}
	return -1
}
