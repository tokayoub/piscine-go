package piscine

func Join(strs []string, sep string) string {
	var strelem string
	for index, v := range strs {
		if index == 0 {
			strelem = strelem + v
		} else {
			strelem = strelem + sep + v
		}
	}
	return strelem
}
