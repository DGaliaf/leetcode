package easy

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	start, end := 0, n

	badVersion := 0
	for start <= end {
		mid := (start + end) / 2

		if isBadVersion(mid) {
			end = mid - 1
			badVersion = mid
		} else {
			start = mid + 1
		}
	}

	return badVersion
}
