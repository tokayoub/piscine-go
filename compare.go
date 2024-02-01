package piscine

func Compare(a, b string) int {
	lena := len(a)
	lenb := len(b)
	for i, n := 0, 0; i < lena && n < lenb; i, n = i+1, n+1 {
		if a[i] < b[n] {
			return -1
		} else if a[i] > b[n] {
			return 1
		}
	}
	if lena > lenb {
		return 1
	} else if lena < lenb {
		return -1
	}
	return 0
}
