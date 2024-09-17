package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(strongPasswordCheckerII("IloveLe3tcode!"))
}

func strongPasswordCheckerII(password string) bool {

	if len(password) < 8 {
		return false
	}

	var (
		isUpper           bool
		isLower           bool
		hasDigit          bool
		hasSpecialChar    bool
		specialCharacters = "!@#$%^&*()-+"
	)

	for i, char := range password {
		if unicode.IsLower(char) {
			isLower = true
		}
		if unicode.IsUpper(char) {
			isUpper = true
		}

		if i != len(password)-1 {
			if password[i] == password[i+1] {
				return false
			}
		}

		if char >= '0' && char <= '9' {
			hasDigit = true
		}

		if contains(specialCharacters, char) {
			hasSpecialChar = true
		}
	}

	if isLower && isUpper && hasSpecialChar && hasDigit {
		return true
	}

	return false
}

func contains(s string, element rune) bool {
	for _, el := range s {
		if el == element {
			return true
		}
	}
	return false
}

/*
It has at least 8 characters. done
It contains at least one lowercase letter.
It contains at least one uppercase letter.
It contains at least one digit.
It contains at least one special character. The special characters are the characters in the following string: "!@#$%^&*()-+".
It does not contain 2 of the same character in adjacent positions (i.e., "aab" violates this condition, but "aba" does not).
*/
