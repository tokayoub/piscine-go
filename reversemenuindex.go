package piscine

func ReverseMenuIndex(menu []string) []string {
	str := make([]string, len(menu))
	for i, n := 0, len(menu)-1; n >= 0; i, n = i+1, n-1 {
		str[i] = menu[n]
	}
	return str
}
