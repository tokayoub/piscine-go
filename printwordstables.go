package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	var b string
	for _, v := range a {
		b = v
		G := []rune(b)
		for _, n := range G {
			z01.PrintRune(n)
		}
		z01.PrintRune('\n')
	}
}
