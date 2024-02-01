package piscine

func BasicJoin(elems []string) string {
	var strelem string
	for _, v := range elems {
		strelem = strelem + v
	}
	return strelem
}
