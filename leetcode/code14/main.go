package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"a"}))
}

func longestCommonPrefix(strs []string) string {
	var output string
	for k := 0; k < len(strs[0]); k++ {
		for i := 1; i < len(strs); i++ {
			if k >= len(strs[i]) || strs[i][k] != strs[i-1][k] {
				return output
			}
		}
		output += string(strs[0][k])
	}
	if len(output) > 0 {
		return output
	}
	return ""
}