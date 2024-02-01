package piscine

func Abort(a, b, c, d, e int) int {
	var str []int
	str = append(str, a, b, c, d, e)
	for i := 0; i < len(str)-1; i++ {
		for n := i + 1; n < len(str); n++ {
			if str[i] > str[n] {
				str[i], str[n] = str[n], str[i]
			}
		}
	}
	return str[2]
}
