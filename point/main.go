package main

import (
"github.com/01-edu/z01"
)
type point struct{
	x int
	y int
}

func PrintInt(a int) {
	r := '0'
	if a/10 > 0 {
		PrintInt(a / 10)
	}
	for i := 0; i < a%10; i++ {
		r++
	}
	z01.PrintRune(r)
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := point{}

	setPoint(&points)
	PrintStr("x = ")
	PrintInt(points.x)
	PrintStr(", y = ")
	PrintInt(points.y)

	z01.PrintRune('\n')
}
