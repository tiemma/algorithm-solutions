package main

import "fmt"

// https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/529/week-2/3293/

/**
 * Definition for a binary tree node.
  * type TreeNode struct {
		 *     Val int
		  *     Left *TreeNode
			 *     Right *TreeNode
			  * }
*/

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func treeDepth(root *TreeNode, count *int) int {
	if root == nil {
		return 0
	}
	len_left := treeDepth(root.Left, count)
	len_right := treeDepth(root.Right, count)
	fmt.Println(len_left, len_right, *count)
	*count = max(len_left+len_right+1, *count)
	return max(len_left, len_right) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	count := 1
	if root == nil {
		return 0
	}
	treeDepth(root, &count)
	return count - 1
}
