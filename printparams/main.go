package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for index, arg := range os.Args {
		if index == 0 {
			continue
		}
		arg1 := []rune(arg)
		for i := 0; i < len(arg1); i++ {
			if index == 0 && (arg1[i] == '.' && arg1[i+1] == '/') {
				i++
			} else {
				z01.PrintRune(arg1[i])
			}
		}
		z01.PrintRune('\n')
	}
}
