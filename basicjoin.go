package challenge_go

func BasicJoin(strs []string) (s string) {
	for _, word := range strs {
		s += word
	}
	return
}
