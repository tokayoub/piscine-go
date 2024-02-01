package piscine

func LastRune(s string) rune {
	leen := len(s)
	G := []rune(s)
	return G[leen-1]
}
