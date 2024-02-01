package piscine

func Rot14(s string) string {
	G := []rune(s)
	for i := 0; i < len(s); i++ {
		if (G[i] >= 'a' && G[i] <= 'l') || (G[i] >= 'A' && G[i] <= 'L') {
			G[i] += 14
		} else if (G[i] >= 'm' && G[i] <= 'z') || (G[i] >= 'M' && G[i] <= 'Z') {
			G[i] -= 12
		}
	}
	return string(G)
}
