package validnumber

func isNumber(s string) bool {

	isDecimal := false
	isExponent := false
	hasSign := false
	hasDigits := false

	// Iterate over characterse in string
	for i := 0; i < len(s); i++ {

		// Check if character is a digit
		if isDigit(s[i]) {
			hasDigits = true
			continue
		}

		// Check if character is a decimal point
		if isDecimalPoint(s[i]) {
			if isDecimal || isExponent {
				return false
			}
			isDecimal = true
			continue
		}

		// Check if character is an exponent
		if isExponentChar(s[i]) {
			if isExponent || !hasDigits {
				return false
			}
			isExponent = true
			// Reset sign flag for exponent sign
			hasSign = false
			// Exponent must be followed by digits
			hasDigits = false
			continue
		}

		// Check if character is a sign
		if isSign(s[i]) {
			if hasSign || hasDigits || isDecimal {
				return false
			}
			hasSign = true
			continue
		}

		// If character is not a digit, decimal point, exponent or sign, return false
		return false

	}

	// Must have at least one digit
	if !hasDigits {
		return false
	}

	return true

}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isDecimalPoint(c byte) bool {
	return c == '.'
}

func isSign(c byte) bool {
	return c == '+' || c == '-'
}

func isExponentChar(c byte) bool {
	return c == 'e' || c == 'E'
}

func main() {
	result := isNumber("-1E+3")
	println(result)
}
