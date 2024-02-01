package piscine

func Compact(ptr *[]string) int {
	i := 0
	q := *ptr
	var str []string
	for _, v := range q {
		if v != "" {
			str = append(str, v)
			i++
		}
	}
	*ptr = str
	return i
}
