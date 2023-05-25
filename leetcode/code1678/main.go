package main

import "fmt"

func main() {
	fmt.Println(interpret("(al)G(al)()()G"))
}

func interpret(command string) string {
	var output string
	for i := 0; i<len(command); i++{
		if string(command[i]) == "G" {
			output += "G"
		}
		if string(command[i]) == "(" && string(command[i+1]) == ")"  {
			output += "o"
		}
		if string(command[i]) == "(" &&  string(command[i+1]) == "a" && string(command[i+2]) == "l" && string(command[i+3]) == ")" {
			output += "al"
		}
	}
    return output
}