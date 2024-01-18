package main

import (
	"fmt"
	"sort"
)

type Queue struct {
	items []int
}
type InputRestrictedQueue struct {
	items []int
}

func (q *InputRestrictedQueue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *InputRestrictedQueue) DequeueFront() int {
	if len(q.items) == 0 {
		panic("queue is empty")
	}
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

// DequeueRear removes a value at the end and returns the removed value
func (q *InputRestrictedQueue) DequeueRear() int {
	if len(q.items) == 0 {
		panic("queue is empty")
	}
	lastIndex := len(q.items) - 1
	toRemove := q.items[lastIndex]
	q.items = q.items[:lastIndex]
	return toRemove
}

// Enqueue adds a value at the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue removes a value at the front
// and RETURNs the removed value

func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

type OutputRestrictedQueue struct {
	items []int
}

// EnqueueRear adds a value at the end
func (q *OutputRestrictedQueue) EnqueueRear(i int) {
	q.items = append(q.items, i)
}

// EnqueueFront adds a value at the front
func (q *OutputRestrictedQueue) EnqueueFront(i int) {
	q.items = append([]int{i}, q.items...)
}

// Dequeue removes a value at the front and returns the removed value
func (q *OutputRestrictedQueue) Dequeue() int {
	if len(q.items) == 0 {
		panic("queue is empty")
	}
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

type PriorityQueue struct {
	items []int
	desc  bool // Determines the type of priority queue: ascending or descending
}

// Enqueue adds an element to the queue
func (pq *PriorityQueue) Enqueue(item int) {
	pq.items = append(pq.items, item)
	if pq.desc {
		sort.Sort(sort.Reverse(sort.IntSlice(pq.items)))
	} else {
		sort.Ints(pq.items)
	}
}

// Dequeue removes and returns the element from the front of the queue
func (pq *PriorityQueue) Dequeue() int {
	if len(pq.items) == 0 {
		panic("priority queue is empty")
	}
	front := pq.items[0]
	pq.items = pq.items[1:]
	return front
}

func main() {

	myQueue := Queue{}
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)

	myQueue1 := InputRestrictedQueue{}
	myQueue1.Enqueue(100)
	myQueue1.Enqueue(200)
	myQueue1.Enqueue(300)
	fmt.Println(myQueue1)

	removedFront := myQueue1.DequeueFront()
	fmt.Println("Removed from front:", removedFront)
	fmt.Println(myQueue1)

	removedRear := myQueue1.DequeueRear()
	fmt.Println("Removed from rear:", removedRear)
	fmt.Println(myQueue1)

	OutputRestrictedQueue := OutputRestrictedQueue{}
	OutputRestrictedQueue.EnqueueRear(100)
	OutputRestrictedQueue.EnqueueRear(200)
	OutputRestrictedQueue.EnqueueFront(300)
	fmt.Println(OutputRestrictedQueue)

	removed := OutputRestrictedQueue.Dequeue()
	fmt.Println("Removed from front:", removed)
	fmt.Println(OutputRestrictedQueue)


	ascQueue := PriorityQueue{desc: false}
    ascQueue.Enqueue(3)
    ascQueue.Enqueue(1)
    ascQueue.Enqueue(2)
    fmt.Println("Ascending Priority Queue:", ascQueue)
    fmt.Println("Dequeue from Ascending:", ascQueue.Dequeue())

    descQueue := PriorityQueue{desc: true}
    descQueue.Enqueue(3)
    descQueue.Enqueue(1)
    descQueue.Enqueue(2)
    fmt.Println("Descending Priority Queue:", descQueue)
    fmt.Println("Dequeue from Descending:", descQueue.Dequeue())
}
