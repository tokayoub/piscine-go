package piscine

func Index(s string, toFind string) int {
	lens := len(s)
	lent := len(toFind)

	if lent == 0 || lens == 0 {
		return 0
	}

	for i := 0; i <= lens-lent; i++ {
		for n := 0; n < lent; n++ {
			if s[i+n] != toFind[n] {
				break
			}
			if n == lent-1 {
				return i
			}
		}
	}

	return -1
}
