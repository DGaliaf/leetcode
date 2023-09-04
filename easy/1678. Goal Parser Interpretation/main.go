package easy

func interpret(command string) string {
	output := ""

	temp := ""
	for _, c := range command {
		if c == 'G' {
			output += "G"
		} else {
			temp += string(c)

			if temp == "()" {
				output += "o"

				temp = ""
			} else if temp == "(al)" {
				output += "al"

				temp = ""
			}
		}
	}

	return output
}
