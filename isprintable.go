package challenge_go

func IsPrintable(s string) bool {
	if StrLen(s) > 0 {
		for _, l := range s {
			if l > 0 && l < 32 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
