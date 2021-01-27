package challenge_go

func BasicAtoi2(s string) int {

	n := 0
	for _, ch := range s {
		if ch > 47 && ch < 58 {
			a := 0
			for b := '1'; b <= ch; b++ {
				a++
			}
			n = n*10 + a
		} else {
			return 0
		}
	}
	return n
}
