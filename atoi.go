package challenge_go

func Atoi(s string) int {
	str := 0
	signCount := 0
	f := false
	for i, v := range s {
		if int(rune(v)) >= 48 && int(rune(v)) <= 57 || v == '+' || v == '-' {
			if (v == '-' || v == '+') && i > 0 {
				return 0
			}
			if v == '-' && i == 0 {
				f = true
			}
			if v == '+' || v == '-' {
				signCount++
				if signCount > 1 {
					return 0
				}
				continue
			}
			a := 0
			for i := '1'; i <= v; i++ {
				a++
			}
			str = str*10 + a
		} else {
			return 0
		}
	}
	if f {
		str = str - (str * 2)
	}
	return str
}
