package piscine

func SortIntegerTable(table []int) {
	len := 0
	for range table {
		len++
	}
	for i := 0; i < len-1; i++ {
		for n := i + 1; n < len; n++ {
			if table[i] > table[n] {
				table[i], table[n] = table[n], table[i]
			}
		}
	}
}
