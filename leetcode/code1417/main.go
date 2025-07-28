package main

import "fmt"

func main() {
	fmt.Println(reformat("a"))
}

func reformat(s string) string {

	var letterSlice, numberSlice []byte

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			numberSlice = append(numberSlice, s[i])
		} else {
			letterSlice = append(letterSlice, s[i])
		}
	}

	var (
		lenLetter = len(letterSlice)
		lenNumber = len(numberSlice)
	)

	if lenLetter == 0 && lenNumber == 1 {
		return string(numberSlice)
	}

	if lenNumber == 0 && lenLetter == 1 {
		return string(letterSlice)
	}

	if lenLetter == 0 || lenNumber == 0 {
		return ""
	}

	

	difference := lenLetter - lenNumber

	if difference != -1 && difference != 0 && difference != 1 {
		return ""
	}

	var output string

	if difference < 0 {
		for i := 0; i < len(numberSlice)-1; i++ {
			output += string(numberSlice[i])
			output += string(letterSlice[i])
			if i == len(numberSlice)-2 {
				output += string(numberSlice[i+1])
			}
		}
	} else if difference > 0 {
		for i := 0; i < len(letterSlice)-1; i++ {
			output += string(letterSlice[i])
			output += string(numberSlice[i])
			if i == len(letterSlice)-2 {
				output += string(letterSlice[i+1])
			}
		}
	} else {
		for i := 0; i < len(letterSlice); i++ {
			output += string(letterSlice[i])
			output += string(numberSlice[i])
		}
	}

	return output
}
