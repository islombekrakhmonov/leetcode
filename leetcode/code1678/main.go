package main

import "fmt"

func main() {
	fmt.Println(interpret("(al)G(al)()()G"))
}

func interpret(command string) string {
	var output string
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			output += "G"
		}
		if command[i] == '(' && (command[i+1]) == ')' {
			output += "o"
		}
		if (command[i]) == '(' && (command[i+1]) == 'a' && (command[i+2]) == 'l' && (command[i+3]) == ')' {
			output += "al"
		}
	}
	return output
}
