package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := 1; i < len(os.Args)-1; i++ {
		for x := i + 1; x < len(os.Args); x++ {
			if os.Args[i] > os.Args[x] {
				os.Args[i], os.Args[x] = os.Args[x], os.Args[i]
			}
		}
	}
	for j := 1; j < len(os.Args); j++ {
		arg := []rune(os.Args[j])
		for n := 0; n < len(arg); n++ {
			z01.PrintRune(arg[n])
		}
		z01.PrintRune('\n')
	}
}
