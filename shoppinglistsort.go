package piscine

func ShoppingListSort(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		for n := i + 1; n < len(slice); n++ {
			if len(slice[i]) > len(slice[n]) {
				slice[i], slice[n] = slice[n], slice[i]
			}
		}
	}
	return slice
}
