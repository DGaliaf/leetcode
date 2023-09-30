package medium

func max(a, b *int) *int {
	if *a > *b {
		return a
	}

	return b
}

func memLeak(memory1 int, memory2 int) []int {
	i := 1
	for ; ; i++ {
		if memory2 < i && memory1 < i {
			break
		}

		if memory2 >= i && memory1 >= i {
			m := max(&memory2, &memory1)
			*m -= i
		} else if memory1 >= i {
			memory1 -= i
		} else {
			memory2 -= i
		}
	}

	return []int{i, memory1, memory2}
}
