package easy

func judgeCircle(moves string) bool {
	var vertical, horizontal int

	for _, move := range moves {
		if move == 'U' {
			vertical++
		}

		if move == 'D' {
			vertical--
		}

		if move == 'L' {
			horizontal++
		}

		if move == 'R' {
			horizontal--
		}
	}

	return horizontal == 0 && vertical == 0
}
