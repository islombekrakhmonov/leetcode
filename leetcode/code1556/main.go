package main

import "fmt"

func main() {
	fmt.Println(thousandSeparator(123456))
}

func thousandSeparator(n int) string {

	output := ""
	nStr := fmt.Sprint(n)

	if len(nStr) <= 3 {
		return nStr
	}

	for i := len(nStr) - 1; i >= 0; i -= 3 {
		start := i - 2
		if start < 0 {
			start = 0
		}
		chunk := nStr[start : i+1]
		if len(chunk) == 3 && start != 0 {
			chunk = "." + chunk
		}
		output = chunk + output
	}

	return output
}
