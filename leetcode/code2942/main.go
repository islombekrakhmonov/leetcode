package main

func main() {
	findWordsContaining([]string{"leet", "code"}, 'e')
}

func findWordsContaining(words []string, x byte) []int {
	var output []int
	for i, v := range words {
		if contains(v, x) {
			output = append(output, i)
		}
	}

	return output
}

func contains(s string, b byte) bool {
	for _, value := range s {
		if value == rune(b) {
			return true
		}
	}

	return false
}
