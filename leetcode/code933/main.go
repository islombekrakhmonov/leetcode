package main

import "fmt"

func main() {
	obj := Constructor()
	param_1 := obj.Ping(1)
	param_2 := obj.Ping(100)
	param_3 := obj.Ping(3001)
	param_4 := obj.Ping(3002)

	fmt.Println(param_1, param_2, param_3, param_4)
}

type RecentCounter struct {
	t []int // timestamps
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.t = append(this.t, t)
	for this.t[0] < t-3000 {
		this.t = this.t[1:]
	}
	return len(this.t)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
