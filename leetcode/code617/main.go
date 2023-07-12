package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func main() {
	// Create nodes
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 3}
	root1.Right = &TreeNode{Val: 2}
	root1.Left.Left = &TreeNode{Val: 5}

	root2 := &TreeNode{Val: 2}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 3}
	root2.Left.Left = nil 
	root2.Left.Right = &TreeNode{Val: 4}
	root2.Right.Left = nil
	root2.Right.Right = &TreeNode{Val: 7}

	fmt.Println(mergeTrees(root1, root2))

}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	head := root1.Val + root2.Val
	root3 := &TreeNode{Val: head}
    
	return root3
}