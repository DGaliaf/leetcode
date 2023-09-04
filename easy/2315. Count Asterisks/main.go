package easy

func countAsterisks(s string) int {
	totalAsteriks := 0
	excludedAsteriks := 0
	barCount := 0

	for _, r := range s {
		if r == '|' {
			barCount++
		}

		if r == '*' {
			totalAsteriks++

			if barCount%2 != 0 {
				excludedAsteriks++
			}
		}
	}

	return totalAsteriks - excludedAsteriks
}
