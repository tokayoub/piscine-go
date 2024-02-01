package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	for i := nb / 2; i > 1; i-- {
		if i*i == nb {
			return i
		}
	}
	return 0
}
