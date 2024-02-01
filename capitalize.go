package piscine

func Capitalize(s string) string {
	G := []rune(s)
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if G[i] >= 'A' && G[i] <= 'Z' {
				i++
			} else if G[i] >= 'a' && G[i] <= 'z' {
				G[i] = G[i] - 32
			}
		}
		if (G[i] >= 'a' && G[i] <= 'z') && i != 0 {
			if (G[i-1] < 'a' || G[i-1] > 'z') && (G[i-1] < 'A' || G[i-1] > 'Z') && (G[i-1] < '0' || G[i-1] > '9') {
				G[i] = G[i] - 32
			}
		} else if (G[i] >= 'A' && G[i] <= 'Z') && i != 0 {
			if (G[i-1] >= 'A' && G[i-1] <= 'Z') || (G[i-1] >= 'a' && G[i-1] <= 'z') || (G[i-1] >= '0' && G[i-1] <= '9') {
				G[i] = G[i] + 32
			}
		}
	}
	return string(G)
}
