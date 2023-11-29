package main

import "fmt"

func main() {
	fmt.Println(distributeCandies(60,4))
}

func distributeCandies(candies int, num_people int) []int {
    answer := make([]int, num_people)
    left := 1
    index := 0

    for candies > 0 {
        if candies > left {
            answer[index] += left
            candies -= left
        } else {
            answer[index] += candies
            candies = 0
        }

        left++
        index = (index + 1) % num_people
    }

    return answer
}