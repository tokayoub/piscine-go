package piscine

func SplitWhiteSpaces(s string) []string {
	var strR []rune
	var strS []string
	var d string
	for index, v := range s {
		if v != ' ' && v != '\n' && v != '\v' {
			strR = append(strR, v)
		}
		if v == ' ' || v == '\n' || v == '\v' || index == len(s)-1 {
			d = string(strR)
			if d != "" {
				strS = append(strS, d)
				strR = strR[:0]
			}
		}
	}
	return strS
}
