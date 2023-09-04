package easy

func romanToInt(s string) int {
	result := 0

	values := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < len(s)-1; i++ {
		if values[string(s[i])] < values[string(s[i+1])] {
			result -= values[string(s[i])]
		} else {
			result += values[string(s[i])]
		}
	}

	return result + values[string(s[len(s)-1])]
}
