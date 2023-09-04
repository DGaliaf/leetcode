package easy

func checkIfExist(arr []int) bool {
	numbers := make(map[int]struct{})

	for i := 0; i < len(arr); i++ {
		if _, ok := numbers[arr[i]*2]; ok {
			return true
		}

		if arr[i]%2 == 0 {
			if _, ok := numbers[arr[i]/2]; ok {
				return true
			}
		}

		numbers[arr[i]] = struct{}{}
	}

	return false
}
