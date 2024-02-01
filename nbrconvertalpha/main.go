package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	char := 96
	for index, nb := range os.Args {
		if index == 0 {
			continue
		} else if os.Args[1] == "--upper" {
			if index == 1 {
				continue
			} else {
				char = 64
				x := 0
				for _, v := range nb {
					if v >= '0' && v <= '9' {
						x = x*10 + int(v-'0')
					}
				}
				if x >= 1 && x <= 26 {
					char = char + x
					z01.PrintRune(rune(char))
				} else {
					z01.PrintRune(' ')
				}

			}
		} else {
			x := 0
			for _, v := range nb {
				if v >= '0' && v <= '9' {
					x = x*10 + int(v-'0')
				}
			}
			if x >= 1 && x <= 26 {
				char = char + x
				z01.PrintRune(rune(char))
			} else {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
