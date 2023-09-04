package easy

func generate(numRows int) [][]int {
	list := [][]int{{1}}

	if numRows >= 2 {
		list = append(list, []int{1, 1})
	}

	for i := 1; i < numRows-1; i++ {
		temp := []int{list[i][0]}

		for j := 0; j < len(list[len(list)-1])-1; j++ {
			temp = append(temp, list[i][j]+list[i][j+1])
		}
		temp = append(temp, list[i][len(list)-1])

		list = append(list, temp)
	}

	return list
}
