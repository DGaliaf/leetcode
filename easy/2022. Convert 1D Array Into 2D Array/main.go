package easy

func construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return [][]int{}
	}

	result := make([][]int, m)

	result[0] = make([]int, n)

	i := 0
	j := 0
	for _, num := range original {
		if n-j == 0 {
			i++
			result[i] = make([]int, n)

			j = 0
		}

		result[i][j] = num

		j++
	}

	return result
}
