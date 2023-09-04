package easy

func isLowercase(r byte) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}

	return false
}

func isUppercase(r byte) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	}

	return false
}

func isSpecial(r byte) bool {
	if r == '!' ||
		r == '@' ||
		r == '#' ||
		r == '$' ||
		r == '%' ||
		r == '^' ||
		r == '&' ||
		r == '*' ||
		r == '(' ||
		r == ')' ||
		r == '-' ||
		r == '+' {
		return true
	}

	return false
}

func isDigit(r byte) bool {
	if r >= '0' && r <= '9' {
		return true
	}

	return false
}

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}

	var hasDigit, hasLowercase, hasUppercase, hasSpecial bool
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			return false
		}

		if isDigit(password[i]) || isDigit(password[i+1]) {
			hasDigit = true
		}

		if isSpecial(password[i]) || isSpecial(password[i+1]) {
			hasSpecial = true
		}

		if isLowercase(password[i]) || isLowercase(password[i+1]) {
			hasLowercase = true
		}

		if isUppercase(password[i]) || isUppercase(password[i+1]) {
			hasUppercase = true
		}
	}

	return hasDigit && hasLowercase && hasUppercase && hasSpecial
}
