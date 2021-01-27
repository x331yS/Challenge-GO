package challenge_go

func AlphaCount(s string) int {
	count := 0
	for _, letter := range s {
		if letter >= 65 && 90 >= letter || letter >= 97 && 122 >= letter {
			count++
		}
	}
	return count
}
