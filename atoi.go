package piscine

func Atoi(s string) int {
	x := 0
	i := 1
	for index, v := range s {
		if index == 0 && v == '+' {
			v++
		} else if index == 0 && v == '-' {
			i = -i
		} else if v >= '0' && v <= '9' {
			x = x*10 + int(v-'0')
		} else {
			return (0)
		}
	}
	return (x * i)
}
