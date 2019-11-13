package split_a_string_in_balanced_strings

func balancedStringSplit(s string) int {
	var L, R, pars int

	for _, v := range s {
		if string(v) == "L" {
			L++
		}else {
			R++
		}
		if L == R {
			pars++
			L = 0
			R = 0
		}
	}

	return pars
}