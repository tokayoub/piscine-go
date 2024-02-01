package piscine

func Split(s, sep string) []string {
	var str1 string
	var str2 []string
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			if s != "" {
				str2 = append(str2, str1)
				str1 = ""
			}
			i += len(sep) - 1
		} else {
			str1 += string(s[i])
		}
	}
	if s != "" {
		str2 = append(str2, str1)
	}
	return str2
}
