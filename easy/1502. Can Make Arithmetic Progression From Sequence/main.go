package easy

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)

	if len(arr) == 2 {
		return true
	}

	dif := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != dif {
			return false
		}
	}

	return true
}
