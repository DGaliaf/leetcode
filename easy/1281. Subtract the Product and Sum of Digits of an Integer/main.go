package easy

func subtractProductAndSum(n int) int {
	sum := 0
	mult := 1
	for n > 0 {
		sum += n % 10
		mult *= n % 10

		n /= 10
	}

	return mult - sum
}
