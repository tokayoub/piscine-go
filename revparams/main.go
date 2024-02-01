package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		arg := []rune(os.Args[i])
		for n := 0; n < len(arg); n++ {
			z01.PrintRune(arg[n])
		}
		z01.PrintRune('\n')
	}
}
