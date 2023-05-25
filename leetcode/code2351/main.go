package main

import "fmt"

func main() {
	fmt.Println(repeatedCharacter("abccbaacz"))
}

func repeatedCharacter(s string) byte {
	var count int
	var letters []string
	var indexes []int
    for i,v := range s{
		count = 0
		for k,l := range s{
			if i != k && v == l {
				count++
				fmt.Println("count:",count,",letter:", string(l),",index:", k)
			}
			if count >0 && count<=1 {
				for _,j := range letters{					
					if j == string(l) {
						fmt.Println(string(j), string(l))
						continue
					} else {
						letters = append(letters, string(l))
						indexes = append(indexes, k)
					} 
				}
			}
		}
	}
	fmt.Println(letters)
	fmt.Println(indexes)
	return byte('l')
}
//The letter 'c' is the first letter to appear twice, because out of all the letters the index of its second occurrence is the smallest.