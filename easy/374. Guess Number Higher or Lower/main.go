package easy

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	start, end := 0, n

	for start <= end {
		mid := (start + end) / 2

		if guess(mid) == 0 {
			return mid
		}

		if guess(mid) > -1 {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}
