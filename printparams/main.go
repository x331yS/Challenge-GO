package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args
	for i, arg := range a {
		if i != 0 {
			for _, argLetter := range arg {
				z01.PrintRune(argLetter)
			}
			z01.PrintRune('\n')
		}
	}
}
