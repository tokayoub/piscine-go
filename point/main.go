package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func tok(char rune) {
	z01.PrintRune(char)
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	do := "x = "
	i := '0'
	n := '0'
	for j := 0; j < points.x/10; j++ {
		i++
	}
	for j := 0; j < points.x%10; j++ {
		n++
	}
	for _, v := range do {
		tok(v)
	}
	tok(i)
	tok(n)
	i = '0'
	n = '0'
	do = ", y = "
	for j := 0; j < points.y/10; j++ {
		i++
	}
	for j := 0; j < points.y%10; j++ {
		n++
	}
	for _, v := range do {
		tok(v)
	}
	tok(i)
	tok(n)
	tok('\n')
}
