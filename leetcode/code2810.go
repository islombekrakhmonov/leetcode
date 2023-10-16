package main

import "fmt"

func main() {
	fmt.Println(finalString("string"))
}

func finalString(s string) string {
    var output string

	for _,v := range s {
		if string(v) == "i"{
			word := ""
			for i:=len(output)-1;i>=0;i--{
				word += string(output[i])
			}
			output = word
		} else {
			output += string(v)
		}
	}

	return output
}