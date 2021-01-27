package challenge_go

func IsNumeric(s string) bool {
	if StrLen(s) > 0 {
		for _, letter := range s {
			if !(letter > 47 && letter < 58) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
