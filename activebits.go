package piscine

func ActiveBits(n int) int {
	i := 0
	var str []int
	for n != 0 {
		nb := n % 2
		str = append(str, nb)
		n = n / 2
	}
	for _, v := range str {
		if v == 1 {
			i++
		}
	}
	return i
}
