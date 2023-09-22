package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main8() {

	root := &TreeNode{Val: 1}
	root.Left = nil
	right := &TreeNode{Val: 2}
	right.Left = &TreeNode{Val: 3}
	root.Right = right

	fmt.Println(preorderTraversal(root))
}

func preorderTraversal(root *TreeNode) []int {
	queue := []*TreeNode{}

	output := []int{}
	queue = append(queue, root)

	for len(queue) > 0 {

		node := queue[0]

		output = append(output, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		fmt.Println(node.Val)
		queue = queue[1:]
		fmt.Println("Len=", len(queue))
	}
	return output
}
