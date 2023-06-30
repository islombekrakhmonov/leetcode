package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ModifyString(str string) string {
    trimmed := strings.Trim(str, " ")
	var builder strings.Builder
	for _, r := range trimmed {
		if !unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}
	removed := builder.String()
	var output string
	for i:=len(removed)-1; i>=0; i--{
		output += string(removed[i])
	}

	return output
}

func main() {
    fmt.Println(ModifyString("oll123eH56"))
	fmt.Println(countBits(127))
}

func countBits(num uint32) int32 {
    s := fmt.Sprintf("%b", num)
	fmt.Println(s)
	mapping := make(map[string]int32)
	for _,v := range s{
		if string(v) == "1" {
			mapping[string(v)]++
		}
	}
	output := mapping["1"]
    
    
    return output
}