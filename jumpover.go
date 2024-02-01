package piscine

func JumpOver(str string) string {
	N := "\n"
	G := []rune(str)
	var str2 string
	if len(str) < 3 {
		return "\n"
	}
	for i := 2; i < len(str); i += 3 {
		str2 = str2 + string(G[i])
	}
	str2 += N
	return str2
}
