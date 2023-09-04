package easy

func findTheDifference(s string, t string) byte {
	mask := byte(0)

	for i := 0; i < len(s); i++ {
		mask ^= s[i] ^ t[i]
	}

	mask ^= t[len(t)-1]

	return mask
}
