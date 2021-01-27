package challenge_go

func Index(s string, toFind string) int {
	x := StrLen(s)
	Find := StrLen(toFind)

	for l := 0; l <= x-Find; l++ {
		if toFind == s[l:l+Find] {
			return l
		}
	}
	return -1
}