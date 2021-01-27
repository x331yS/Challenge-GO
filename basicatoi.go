package challenge_go

func BasicAtoi(s string) int {
	n := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			return 0
		}
		n = n*10 + int(ch)

	}
	return n
}
