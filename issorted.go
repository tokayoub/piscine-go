package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var i int
	for i, n := 0, i+1; n < len(a); i, n = i+1, n+1 {
		if f(a[i], a[n]) > 0 && n < len(a)-1 {
			if f(a[i+1], a[n+1]) < 0 {
				return false
			}
		} else if f(a[i], a[n]) < 0 && n < len(a)-1 {
			if f(a[i+1], a[n+1]) > 0 {
				return false
			}
		}
	}
	return true
}
