package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Put(2, 3)
	param_2 := obj.Get(2)
	fmt.Println(param_2)
	obj.Remove(2)
}

type MyHashMap struct {
}

func Constructor() MyHashMap {

	return MyHashMap{}
}

func (this *MyHashMap) Put(key int, value int) {

	return
}

func (this *MyHashMap) Get(key int) int {

}

func (this *MyHashMap) Remove(key int) {

}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
