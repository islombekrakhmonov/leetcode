package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
}

func numUniqueEmails(emails []string) int {
	uniqueEmails := make(map[string]int)
	for i := 0; i < len(emails); i++ {
		for j := 0; j < len(emails[i]); j++ {
			if emails[i][j] == '@' {
				break
			}
			if emails[i][j] == '.' {
				emails[i] = emails[i][:j] + emails[i][j+1:]
			}
			if emails[i][j] == '+' {
				emails[i] = removeCharacters(emails[i])
				break
			}

		}
		uniqueEmails[emails[i]]++
	}

	return len(uniqueEmails)
}

func removeCharacters(input string) string {
	dashIndex := strings.Index(input, "+")
	ampersandIndex := strings.Index(input, "@")

	if dashIndex != -1 && (ampersandIndex == -1 || dashIndex < ampersandIndex) {
		result := input[:dashIndex] + input[ampersandIndex:]
		return result
	}

	return input
}
