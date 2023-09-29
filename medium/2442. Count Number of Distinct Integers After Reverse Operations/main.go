package medium

func reverseNum(n int) (res int) {
	for n > 0 {
		remainder := n % 10
		res *= 10
		res += remainder
		n /= 10
	}
	return res
}

func countDistinctIntegers(nums []int) int {
	seen := make(map[int]struct{}, 0)
	distinct := 0

	for _, num := range nums {
		if _, ok := seen[num]; !ok {
			seen[num] = struct{}{}
			distinct++
		}

		reversed := reverseNum(num)

		if _, ok := seen[reversed]; !ok {
			seen[reversed] = struct{}{}
			distinct++
		}
	}

	return distinct
}
