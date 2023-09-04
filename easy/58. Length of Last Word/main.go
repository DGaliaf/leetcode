package easy

func lengthOfLastWord(s string) int {
	res := 0
	count := 0
	for _, r := range s {
		if r != ' ' {
			count++
		} else {
			if count != 0 {
				res = count
			}
			count = 0
		}
	}

	if count != 0 {
		res = count
	}

	return res
}
