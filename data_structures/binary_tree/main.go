package main

import "fmt"

var count int 
type Node struct {
	Key int 
	Right *Node
	Left *Node 
}

//Insert will add a node to the tree

func (n *Node) Insert (k int) {
	if n.Key < k{
		// move left
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move right
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

//Search will take in a key value
// and return true if there is a node with that value 

func (n *Node) Search (k int) bool {
	count++
	if n == nil {
		return false
	}

	if n.Key <k {
		//move right
		return n.Right.Search(k)
	} else if n.Key > k {
		//move left 
		return	n.Left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{Key:100}
	tree.Insert(50)
	tree.Insert(200)
	tree.Insert(400)
	tree.Insert(32)
	tree.Insert(30)
	tree.Insert(240)
	tree.Insert(540)
	tree.Insert(37)
	tree.Insert(50)
	tree.Insert(60)
	tree.Insert(500)

	fmt.Println(tree)
	fmt.Println(tree.Search(37))
	fmt.Println(count)
	
}
