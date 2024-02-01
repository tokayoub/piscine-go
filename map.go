package piscine

func Map(f func(int) bool, a []int) []bool {
	Bo := make([]bool, len(a))
	for i, n := 0, 0; i < len(a); i, n = i+1, n+1 {
		Bo[n] = f(a[i])
	}
	return Bo
}
