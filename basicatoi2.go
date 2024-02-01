package piscine

func BasicAtoi2(s string) int {
	x := 0
	for _, v := range s {
		if v >= '0' && v <= '9' {
			x = x*10 + int(v-'0')
		} else {
			return (0)
		}
	}
	return (x)
}
