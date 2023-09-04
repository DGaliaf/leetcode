package easy

import "math"

func nearestValidPoint(x int, y int, points [][]int) int {
	minManhattanDistance := math.MaxInt
	ans := -1
	for i, point := range points {
		if point[0] == x || point[1] == y {
			manhattanDistance := int(math.Abs(float64(point[0]-x))) + int(math.Abs(float64(point[1]-y)))

			if manhattanDistance < minManhattanDistance {
				minManhattanDistance = manhattanDistance
				ans = i
			}
		}
	}

	return ans
}
