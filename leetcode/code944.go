package main

import "fmt"

func main() {
	fmt.Println(minDeletionSize([]string{"abc", "bce", "cae"}))
}

func minDeletionSize(strs []string) int {
	count := 0
    
    for j := 0; j < len(strs[0]); j++ {
        for i := 1; i < len(strs); i++ {
            if strs[i-1][j] > strs[i][j] {
				fmt.Println(string(strs[i-1]))
				fmt.Println(string(strs[i]))
                count++
                break 
            }
        }
    }

	return count
}

// var sIndex int 
// 	var foundB bool
//     for sIndex < len(s) {
// 		if string(s[sIndex]) == "b"{
// 			foundB = true
// 		}
// 		if string(s[sIndex]) == "a" && foundB {
// 			return false 
// 		}
// 		sIndex++
// 	}