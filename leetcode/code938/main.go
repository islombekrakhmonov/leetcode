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
	root.Right.Left = nil
	root.Right.Right = &TreeNode{Val: 18}

	
	fmt.Println(rangeSumBST(root, 7, 15))

	fmt.Println(root.Right.Val)
	fmt.Println(root)
	fmt.Println(checkTree(root))
}



func rangeSumBST(root *TreeNode, low int, high int) int {

	if root == nil {
		 return 0
	 }
 
	 sum := 0
	 
	 if root.Val >= low && root.Val <= high {
		 sum += root.Val
	 }
 
	 l := rangeSumBST(root.Left, low, high)
 
	 r := rangeSumBST(root.Right, low, high)
	
	 return sum + l + r
 }




func checkTree(root *TreeNode) bool {
    return root.Right.Val + root.Left.Val == root.Val  
}

