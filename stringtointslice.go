package piscine

func StringToIntSlice(str string) []int {
	var str1 []int
	if str == "" {
		return nil
	}
	for _, v := range str {
		str1 = append(str1, int(v))
	}
	return str1
}
