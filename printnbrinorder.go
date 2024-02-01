package piscine

import "github.com/01-edu/z01"

func strsor(s []int) {
	for i := 0; i < len(s)-1; i++ {
		for n := i + 1; n < len(s); n++ {
			if s[i] > s[n] {
				s[i], s[n] = s[n], s[i]
			}
		}
	}
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	var result []int
	for n > 0 {
		digit := n % 10
		result = append(result, digit)
		n = n / 10
	}
	strsor(result)
	for i := 0; i < len(result); i++ {
		z01.PrintRune(rune(result[i]) + '0')
	}
}
