package easy

func hammingDistance(x int, y int) int {
	count := 0

	for i := 0; i < 32; i++ {
		if (x>>i)&1 != (y>>i)&1 {
			count++
		}
	}

	return count
}
