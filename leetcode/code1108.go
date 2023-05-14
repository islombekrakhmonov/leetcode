package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(defangIPaddr("1.1.1.1"))
	
}

func defangIPaddr(address string) string {
	output := strings.Replace(address, ".", "[.]", -1)
	return output
}