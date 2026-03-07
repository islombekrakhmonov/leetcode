package main

import "fmt"

func main() {
	fmt.Println(getMaximumGenerated(7))
}

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	array := make([]int, n+1)
	array[0] = 0
	array[1] = 1

	max := 0
	for i := 1; i <= n/2; i++ {
		if 2*i <= n && 2 <= 2*i {
			array[2*i] = array[i]
			if array[2*i] > max {
				max = array[2*i]
			}
		}
		if (2*i)+1 <= n {
			array[(2*i)+1] = array[i] + array[i+1]
			if array[(2*i)+1] > max {
				max = array[(2*i)+1]
			}
		}
	}
	return max
}

//if index % 2 == 0, the value of current index will be of the same if i/2. And if index % 2! = 0, then the value of current index will be (index -1 / 2) + index next to it.
