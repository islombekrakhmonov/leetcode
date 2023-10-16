package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numOfStrings([]string{"dojsf","v","hetnovaoigv","ksvqfdojsf","hetnovaoig","yskjs","sfr","msurztkvppptsluk","ndxffbkkvejuakduhdcfsdbgbt","znhdgtwzbnh","h","ocaualfiscmbpnfalypmtdppymw","w","o","sfjksvqfdo","odqvzuc","bozawheajcmlewptgssueylcxhx","bno","jhmarzsphxduvdktzqjiz","j","sfrjhetnov","vxv","ksvqfd","ognwvqtadalmav","yqbspvfwmvhycmghabigl","qyfaiazgdqaw","ojsfrj","ojsfrjhetn","fdojs","do","ovaoig","k","vrh","jsfrjhetnovaoigvgk"}, "csfjksvqfdojsfrjhetnovaoigvgk"))
}

func numOfStrings(patterns []string, word string) int {
    var output int
    for _, pattern := range patterns {
        if strings.Contains(word, pattern) {
            output++
        }
    }
    return output
}