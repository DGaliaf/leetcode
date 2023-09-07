package medium

func findArray(pref []int) []int {
	output := make([]int, len(pref))

	output[0] = pref[0]
	for i := 0; i < len(pref)-1; i++ {
		output[i+1] = pref[i] ^ pref[i+1]
	}

	return output
}
