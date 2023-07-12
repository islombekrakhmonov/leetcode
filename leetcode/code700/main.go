package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func main() {
	// Create nodes
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	// root.Right.Left = nil
	// root.Right.Right = &TreeNode{Val: 18}

	fmt.Println(searchBST(root, 2))
	
}


func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
		return nil
	} 

	if root.Val == val {
		return root
	}

	if root.Val < val {
		return searchBST(root.Right, val)
		
	} else if root.Val > val {
		return searchBST(root.Left, val)
	}

	return nil
}
