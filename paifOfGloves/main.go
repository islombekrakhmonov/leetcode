package main

import "fmt"

// input = ["red", "green", "red", "blue", "blue"]
func main(){
	gogo := []string{"red", "green", "red", "blue", "blue","red","red", "green", "blue", "blue"}
	fmt.Println(NumberOfPairs(gogo))
}
func NumberOfPairs(gloves []string) int {
	var output int
	m := make(map[string]int)


	for _,v := range gloves{
		if m[v] !=0{
		m[v] +=1
		} else {
			m[v]=1
		}
		fmt.Println(m)
	} 
	for _,v:=range m{
		if v>1{
		  output+=v/2
		}
	  } 	
	return output
}

