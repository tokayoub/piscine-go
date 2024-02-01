package piscine

func NRune(s string, n int) rune {
	leen := len(s)
	if n < 1 || n > leen {
		return 0
	}
	G := []rune(s)
	return G[n-1]
}
