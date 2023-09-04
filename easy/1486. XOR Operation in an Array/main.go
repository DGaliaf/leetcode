package easy

func xorOperation(n int, start int) (res int) {
	num := start
	for i := 0; i < n; i++ {
		res ^= num
		num += 2
	}

	return
}
