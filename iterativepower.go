package piscine

func IterativePower(nb int, power int) int {
	x := nb
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	for i := 1; i < power; i++ {
		nb = nb * x
	}
	return nb
}
