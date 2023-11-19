package main

import "fmt"

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1,1,0},{1,0,1},{0,0,0}}))
}

func flipAndInvertImage(image [][]int) [][]int {
    for i, v := range image {
		for j, k := range v {
			if k == 0 {
				image[i][j] = 1
			} else {
				image[i][j] = 0
			}
		}
	}
	
	for i:=0; i<len(image); i++ {
		start := 0
    	end := len(image[i]) - 1
		for start < end {
			image[i][start], image[i][end] = image[i][end], image[i][start]
        	start++
        	end--
		}
	}

	return image
}