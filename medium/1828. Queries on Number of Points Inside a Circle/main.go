package medium

func pow(val int) int {
	return val * val
}

func countPoints(points [][]int, queries [][]int) []int {
	output := make([]int, len(queries))

	for i, query := range queries {
		count := 0

		for _, point := range points {
			if pow(point[0]-query[0])+pow(point[1]-query[1]) <= pow(query[2]) {
				count++
			}
		}

		output[i] = count
	}

	return output
}
