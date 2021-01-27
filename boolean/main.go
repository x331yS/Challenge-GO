package main

import("github.com/01-edu/z01"
"os")

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	lengthOfArg :=0
	EvenMsg := "I have an odd number of arguments"
	OddMsg:= "I have an even number of arguments"
	for range os.Args {
		lengthOfArg ++
	}
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}