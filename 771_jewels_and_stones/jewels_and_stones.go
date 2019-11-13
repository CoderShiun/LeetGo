package jewels_and_stones

func numJewelsInStones(J string, S string) int {
	var num int
	for _, s := range S {
		for _, j := range J {
			if j == s {
				num++
				break
			}
		}
	}

	return num
}
