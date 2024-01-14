package main

import "fmt"


type Queue struct{
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
}