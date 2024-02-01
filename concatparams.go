package piscine

func ConcatParams(args []string) string {
	var str string
	for index, v := range args {
		str = str + v
		if index < len(args)-1 {
			str = str + "\n"
		}
	}
	return str
}
