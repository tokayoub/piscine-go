package piscine

func ToLower(s string) string {
	G := []rune(s)
	lens := len(s)
	for i := 0; i < lens; i++ {
		if G[i] >= 'A' && G[i] <= 'Z' {
			G[i] = G[i] + 32
		}
	}
	return string(G)
}
