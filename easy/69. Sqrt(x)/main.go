package easy

import "math"

func mySqrt(x int) int {
	if x < 2 {
		return 1
	}

	y := float64(x)
	d := (y + (float64(x) / y)) / 2
	for math.Abs(y-d) > 0.1 {
		y = d
		d = (y + (float64(x) / y)) / 2
	}

	return int(d)
}
