package main

import "fmt"

func main() {
	fmt.Println(timeRequiredToBuy([]int{84,49,5,24,70,77,87,8}, 3))
}

type Queue struct {
	queue []int
}

func (q *Queue) Enqueue(i int) {
	q.queue = append(q.queue, i)
}

func (q *Queue) Dequeue() {
	q.queue = q.queue[1:]
}

func timeRequiredToBuy(tickets []int, k int) int {
	var totalTime int

	for tickets[k] > 0{
		for i :=0; i<len(tickets); i++ {
			if tickets[i] > 0 {
				tickets[i]--
				totalTime++
			}
			if tickets[k] == 0 {
				return totalTime
			}
		}
	}

	return totalTime
}