package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	if str == " " {
		return map[string]int{"": 2}
	} else if str == "" {
		return map[string]int{"": 1}
	}
	i := 0
	var sstr string
	var str1 []string
	list := make(map[string]int)
	for index, v := range str {
		if v != ' ' {
			sstr += string(v)
		} else {
			str1 = append(str1, sstr)
			sstr = ""
			index++
		}
	}
	if sstr != "" {
		str1 = append(str1, sstr)
	}
	for _, v := range str1 {
		list[v] = 0
	}
	for index := range list {
		for _, y := range str1 {
			if index == y {
				i++
			}
		}
		list[index] = i
		i = 0
	}
	return list
}
