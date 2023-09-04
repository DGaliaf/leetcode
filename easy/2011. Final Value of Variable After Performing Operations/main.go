package easy

func finalValueAfterOperations(operations []string) int {
	res := 0

	for _, operation := range operations {
		if operation[0] == '+' || operation[len(operation)-1] == '+' {
			res++
		}

		if operation[0] == '-' || operation[len(operation)-1] == '-' {
			res--
		}
	}

	return res
}
