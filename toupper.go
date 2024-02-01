package piscine

func ToUpper(s string) string {
	G := []rune(s)
	lens := len(s)
	for i := 0; i < lens; i++ {
		if G[i] >= 'a' && G[i] <= 'z' {
			G[i] = G[i] - 32
		}
	}
	return string(G)
}
