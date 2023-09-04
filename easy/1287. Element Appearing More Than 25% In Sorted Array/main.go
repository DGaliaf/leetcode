package easy

func findSpecialInteger(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	count := 1
	curEl := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] == curEl {
			count++

			if count > len(arr)/4 {
				return curEl
			}
		} else {
			count = 1
			curEl = arr[i]
		}
	}

	return -1
}
