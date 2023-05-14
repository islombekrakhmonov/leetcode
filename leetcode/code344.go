package main

import (
	"fmt"
	// "sort"
)

func main(){
   fmt.Println(reverse([]byte{'h','e','l','l','o'}))
}

func reverseString(s []byte) []byte {
	for i, j := 0, len(s)-1; i<j; {
		s[i],s[j] = s[j],s[i]
	}
	return s
}

// func reverseString1(s []byte)  {
//     sort.SliceStable(s, func(a, b int) bool {
//         return a > b
//     })
// }


func reverse(s []byte)[]string{
    var output []string;
            for i := len(s)-1; i >= 0; i-- {
                output = append(output, string(s[i]))
            }
            return output
}
        


