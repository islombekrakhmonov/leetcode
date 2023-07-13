package main

import (
	"fmt"
)

//ArraySize is the size of the hash table array
const ArraySize = 7
//HashTable structure 

type HashTable struct{
	array [ArraySize]*bucket
}

type bucket struct{
	head *bucketNode
}

type bucketNode struct{
	key interface{}
	next *bucketNode
}


//bucker structure
// bucketNode structure 

//Insert
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}
//Search 
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}
//Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

//insert will take in a key, create a node with the key and insert the node in the bucket 
func (b *bucket) insert(k string){
	if !b.search(k) {
		newNode := &bucketNode{ key: k }
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}
//search
func (b *bucket) search(k string)bool{
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k{
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
//delete
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return 
	}
	previousNode := b.head
	for previousNode.next != nil{
		if previousNode.next.key == k{
			// delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

//hash
func hash(key string) int{
	sum := 0
	for _,v := range key{
		sum +=int(v)
	}
	return sum % ArraySize
}
//Init will create a bucket in each slot of the hash table 
func Init() *HashTable{
	result := &HashTable{}
	for i:= range result.array{
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"JOHN",
		"STAN",
		"RANDY",
		"TOKEN",
		"SARA",
	}

	for _, v:= range list{
		hashTable.Insert(v)
	}

	fmt.Println(hashTable.array)
	fmt.Println(85^182)
}