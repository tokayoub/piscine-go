package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg1 := []rune(os.Args[0])
	for i := 2; i < len(arg1); i++ {
		z01.PrintRune(arg1[i])
	}
	z01.PrintRune('\n')
}
