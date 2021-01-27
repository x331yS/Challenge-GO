package challenge_go

func IsAlpha(s string) bool {
	if StrLen(s) > 0 {
		for _, letter := range s {
			if !(letter > 47 && letter < 58) && !(letter > 64 && letter < 91) && !(letter > 96 && letter < 123) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
