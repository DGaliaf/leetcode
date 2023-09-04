package easy

func toLowerCase(s string) string {
	sBytes := []byte(s)

	for i := 0; i < len(sBytes); i++ {
		if sBytes[i] >= 65 && sBytes[i] <= 90 {
			sBytes[i] = sBytes[i] + 32
		}
	}

	return string(sBytes)
}
