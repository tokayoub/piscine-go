package piscine

func AlphaCount(s string) int {
	n := 0
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			n++
		}
	}
	return n
}
