package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{0, 0, 1, 0, 0}, 1))
}

func canPlaceFlowers(flowerbed []int, n int) bool {

	// for i := 1; i < len(flowerbed)-1; i++ {
	// 	if n > 0 {
	// 		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
	// 			flowerbed[i] = 1
	// 			n--
	// 		}
	// 	}
	// }

	for i := 0; i < len(flowerbed); i++ {
		if n > 0 {
			if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		}
	}

	return n == 0
}
