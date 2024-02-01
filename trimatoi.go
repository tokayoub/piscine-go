package piscine

func TrimAtoi(s string) int {
	x := 0
	i := 1
	for index, v := range s {
		if index == 2 && v == '-' {
			i = -i
		} else if v >= '0' && v <= '9' {
			x = x*10 + int(v-'0')
		}
	}
	return (x * i)
}
